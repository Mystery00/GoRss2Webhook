package handler

import (
	"GoRss2Webhook/core"
	"GoRss2Webhook/feed/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type FeedSubscriberRequest struct {
	FeedUrl   string `json:"feedUrl"  binding:"required"`
	UserAgent string `json:"userAgent"`
	ProxyUrl  string `json:"proxyUrl"`
	Timeout   string `json:"timeout"`
}

func subscribeRssFeed(context *gin.Context) {
	var request FeedSubscriberRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		logrus.Warn(err)
		context.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	duration, err := time.ParseDuration(request.Timeout)
	if err != nil {
		duration = 0
	}
	subscriber := store.FeedSubscriber{
		FeedUrl:   request.FeedUrl,
		UserAgent: request.UserAgent,
		ProxyUrl:  request.ProxyUrl,
		Timeout:   duration,
	}
	s := core.FeedStore()
	err = (*s).Subscribe(subscriber)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"success": true,
	})
}

func unsubscribeRssFeed(context *gin.Context) {
	var request FeedSubscriberRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		logrus.Warn(err)
		context.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	s := core.FeedStore()
	err = (*s).Unsubscribe(request.FeedUrl)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"success": true,
	})
}
