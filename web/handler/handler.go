package handler

import (
	"GoRss2Webhook/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handle(engine *gin.Engine) {
	//注册路由
	engine.GET(apiPath+"/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	//Rss订阅
	engine.POST(externalPath+"/feed", subscribeRssFeed)
	engine.DELETE(externalPath+"/feed", unsubscribeRssFeed)
	//Webhook
	engine.POST(externalPath+"/webhook", newWebhook)
	engine.DELETE(externalPath+"/webhook", deleteWebhook)
	//触发器
	engine.PUT(externalPath+"/check", func(context *gin.Context) {
		core.DoWork()
		context.JSON(204, nil)
	})
}
