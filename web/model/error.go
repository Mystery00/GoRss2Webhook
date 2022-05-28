package model

import "github.com/gin-gonic/gin"

type ErrorMessage struct {
	StatusCode int
	Code       int
	Message    string
}

func (message ErrorMessage) Json() gin.H {
	return gin.H{
		"code": message.Code,
		"msg":  message.Message,
	}
}
