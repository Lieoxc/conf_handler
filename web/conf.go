package web

import (
	"cfg_handler/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func CfgResp(cfg interface{}, codeMsg model.CodeMsg) *model.WebResp {
	return &model.WebResp{
		CodeMsg: codeMsg,
		Data:    cfg,
	}
}

func CfgGetHandler(c *gin.Context) {
	// 1. 解析请求参数
	token := c.Query("token")
	// 2. 解析本地的 subinfo
	addrs, err := model.GetSubProxies(token)
	if len(addrs) == 0 {
		c.JSON(http.StatusOK, CfgResp(nil, model.NewCodeMsg(-1, "订阅地址非法")))
	}
	addMap := make(map[string]struct{}, 0)
	for _, addr := range addrs {
		key := addr.Server + ":" + addr.InnerName
		addMap[key] = struct{}{}
	}
	// token 参数有效
	tempCfg := model.GlobalCfg
	// 处理 clash 的 proxies 配置，可选的服务器只能是select中允许的范围内
	tempCfg.Proxies = []model.Proxy{} // 清空配置

	allProxies, err := model.GetProxies()

	var proxies []model.Proxy
	for _, proxieSt := range allProxies {
		var proxie model.Proxy
		//fmt.Println("conf:", proxieSt.Conf)
		err = json.Unmarshal([]byte(proxieSt.Conf), &proxie)
		if err != nil {
			fmt.Println("json Unmarshal conf error")
		}
		key := proxieSt.Server + ":" + proxieSt.InnerName
		if _, ok := addMap[key]; ok {
			proxie.Name = proxieSt.Name
			proxies = append(proxies, proxie)
		}
	}
	tempCfg.Proxies = proxies

	serverNames, err := model.GetNamesFromAddrs(addrs)
	if err != nil {
		fmt.Println("GetNamesFromAddrs err:", err)
		return
	}

	// 处理 clash 的 proxy-groups 配置, 可选的服务器只能是select中允许的范围内
	for i, group := range tempCfg.ProxyGroups {
		newGroup := group
		newGroup.Proxies = serverNames
		tempCfg.ProxyGroups[i] = newGroup
	}

	newRules, err := model.GetCustomRule()
	if err != nil {
		fmt.Println("GetCustomRule err:", err)
		return
	}

	newRules = append(newRules, tempCfg.Rules...)
	tempCfg.Rules = newRules
	cfgData, err := yaml.Marshal(&tempCfg)
	if err != nil {
		fmt.Println("Marshal allProxy err", err)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=UTF-8", cfgData)
}

func DownloadUrlHandler(c *gin.Context) {
	url, err := model.GetDownloadUrl()
	if err != nil {
		fmt.Println("GetNamesFromAddrs err:", err)
		return
	}
	fmt.Println(url)
	c.JSON(200, Reply{
		Code: 0,
		Data: url,
		Msg:  "操作成功",
	})
}
