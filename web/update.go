package web

import (
	"cfg_handler/model"
	"cfg_handler/utils"
	"encoding/json"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func SubUpdateHandler(c *gin.Context) {
	utils.UpdateData()
	c.JSON(200, Reply{
		Code: 0,
		Msg:  "操作成功",
	})
}

type UpdateRequest struct {
	ID int `json:"id"` // 注意字段名和 JSON 标签要与请求体一致
}

func ProxieUpdateHandler(c *gin.Context) {
	id := c.Query("id")
	rawData, err := model.GetSubRaw(id)
	if err != nil {
		c.JSON(200, Reply{
			Code: 1,
			Msg:  "操作失败",
		})
		return
	}
	var allProxy []model.Proxy
	for _, data := range rawData {
		// 序列化到一个结构体
		var cfg model.OpenclashCfg
		// 解析YAML数据
		err = yaml.Unmarshal([]byte(data.RawData), &cfg)
		if err != nil {
			log.Printf("Error parsing YAML file: %v", err)
			continue
		}
		for i := range cfg.Proxies {
			cfg.Proxies[i].SubFrom = data.RawUrl
		}
		// 保存到所有服务器节点
		allProxy = append(allProxy, cfg.Proxies...)
	}
	// 更新ProxyMap
	words, err := model.GetWords()
	if err != nil {
		log.Printf("GetWords failed %v", err)
		return
	}
	wordsMap := make(map[string]struct{}, len(words))
	for _, word := range words {
		wordsMap[word.Word] = struct{}{}
	}
	proxyMap := model.UpdateProxyMap(allProxy, wordsMap)

	// 保存到数据库Proxies表
	for key, proxies := range proxyMap {
		keys := strings.Split(key, ":") // keys[0] - serverAddr ;  keys[1] - name
		confBytes, err := json.Marshal(proxies)
		if err != nil {
			log.Printf("json.Marshal proxies : %v", err)
			continue
		}
		proxie := model.Proxies{
			Name:       keys[1],
			InnerName:  keys[1],
			Server:     proxies.Server,
			Conf:       string(confBytes),
			ProxieType: "auto",
			FromSub:    proxies.SubFrom,
		}
		exist, err := model.CheckExist(proxies.Server, proxie.Name)
		if !exist {
			model.CreateProxie(proxie)
		} else {
			model.UpdateProxie(proxie)
		}
	}
	c.JSON(200, Reply{
		Code: 0,
		Msg:  "操作成功",
	})
}
