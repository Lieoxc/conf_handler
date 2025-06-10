package web

import "github.com/gin-gonic/gin"

func WebStart() {

	r := gin.Default()
	api := r.Group("/api/")
	api.GET("/conf/", CfgGetHandler)              // 获取openclash的配置信息，全量
	api.GET("/downloadUrl", DownloadUrlHandler)   // 获取clash内核的下载地址
	api.POST("/subUpdate", SubUpdateHandler)      // 立即更新订阅
	api.GET("/proxieUpdate", ProxieUpdateHandler) // 手动解析节点
	r.Run(":18083")

}
