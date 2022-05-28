package handler

import (
	"GoRss2Webhook/core"
	"GoRss2Webhook/webhook/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type WebhookRequest struct {
	SubscribeUrl string `json:"subscribeUrl"  binding:"required"`
	Type         int8   `json:"type"  binding:"required"`
	MetaData     string `json:"metaData"  binding:"required"`
}

func newWebhook(context *gin.Context) {
	feedType := viper.GetString(core.StoreWebhookType)
	if feedType == "viper" {
		context.JSON(400, gin.H{
			"code": 400,
			"msg":  "error webhook type",
		})
		return
	}
	var request WebhookRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		logrus.Warn(err)
		context.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	s := core.WebhookStore()
	err = (*s).Save(request.SubscribeUrl, store.Webhook{
		Type:     request.Type,
		MetaData: request.MetaData,
	})
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"success": true,
	})
}

type DeleteWebhookRequest struct {
	SubscribeUrl string `json:"subscribeUrl"  binding:"required"`
}

func deleteWebhook(context *gin.Context) {
	feedType := viper.GetString(core.StoreWebhookType)
	if feedType == "viper" {
		context.JSON(400, gin.H{
			"code": 400,
			"msg":  "error webhook type",
		})
		return
	}
	var request DeleteWebhookRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		logrus.Warn(err)
		context.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	s := core.WebhookStore()
	err = (*s).Delete(request.SubscribeUrl)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"success": true,
	})
}
