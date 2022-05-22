package handler

import (
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
}
