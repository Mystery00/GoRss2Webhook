package middleware

import (
	"GoRss2Webhook/web/model"
	"github.com/gin-gonic/gin"
)

var noRouteMiddleware gin.HandlerFunc = func(context *gin.Context) {
	panic(model.ErrorMessage{
		Code:    404,
		Message: "Not Found",
	})
}
