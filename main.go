package main

import (
	"cfg_handler/conf"
	"cfg_handler/model"
	"cfg_handler/utils"
	"cfg_handler/web"
	"time"
)

func main() {

	conf.Init()
	utils.InitGlobalCfg()
	model.InitDb()
	go loopRun()
	// 开启一个web服务
	web.WebStart()
}
func loopRun() {
	utils.UpdateData()

	ticker := time.NewTicker(1 * time.Hour)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				utils.UpdateData()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// 保持程序持续运行
	select {}
}
