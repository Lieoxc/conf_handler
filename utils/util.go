package utils

import (
	"cfg_handler/model"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

func UpdateData() {
	urls, err := model.GetUrls()
	if err != nil {
		fmt.Println("InitUrls err", err)
		return
	}
	//initGlobalFlag := false
	var allProxy []model.Proxy
	for i, val := range urls {
		log.Printf("url: %v", val.Urls)
		fileName := fmt.Sprintf("./data/subInfo_%d.yaml", i)
		cfgVal, err := HTTPGet(val.Urls, fileName)
		if err != nil {
			fmt.Println("HTTPGet err", err)
			err = model.UpdateUrlInfo(val.Urls, model.StatusExpect, 0)
			if err != nil {
				log.Printf("UpdateUrlInfo failed %s", val.Urls)
			}
			continue
		}
		// 序列化到一个结构体
		var cfg model.OpenclashCfg
		// 解析YAML数据
		err = yaml.Unmarshal(cfgVal, &cfg)
		if err != nil {
			log.Printf("Error parsing YAML file: %v", err)
			err = model.UpdateUrlInfo(val.Urls, model.StatusExpect, 0)
			if err != nil {
				log.Printf("UpdateUrlInfo failed %s", val.Urls)
			}
			continue
		}
		for i := range cfg.Proxies {
			cfg.Proxies[i].SubFrom = val.Urls
		}
		err = model.UpdateUrlInfo(val.Urls, model.StatusOk, len(cfg.Proxies))
		if err != nil {
			log.Printf("UpdateUrlInfo failed %s", val.Urls)
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
}
func HTTPGet(urlStr, fileName string) ([]byte, error) {
	// 解析 URL
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		log.Printf("Error creating request for %s: %v", urlStr, err)
		return []byte{}, err
	}
	// HTTP客户端
	client := &http.Client{
		Timeout: 10 * time.Second, // 设置超时时间
	}
	// 设置User-Agent头
	req.Header.Set("User-Agent", "clash")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request to %s: %v", urlStr, err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response from %s: %v", urlStr, err)
		return []byte{}, err
	}
	// 保存到一个文件
	err = ioutil.WriteFile(fileName, body, 0644)
	if err != nil {
		log.Fatalf("Error writing YAML file: %v", err)
	}
	return body, nil
}

func InitGlobalCfg() error {
	// 读取 YAML 格式的配置文件内容
	yamlContent, err := ioutil.ReadFile("./conf/model.yaml")
	if err != nil {
		log.Printf("failed to read config.yml:%v", err)
		return err
	}

	var cfg model.OpenclashCfg
	// 解析YAML数据
	err = yaml.Unmarshal(yamlContent, &cfg)
	if err != nil {
		log.Printf("Error parsing YAML file: %v", err)
		return err
	}
	model.GlobalCfg = cfg
	log.Printf("xxxxxxxxxxxx GlobalCfg:%d", model.GlobalCfg.MixedPort)
	return nil
}
