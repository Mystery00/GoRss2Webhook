package main

import (
	"GoRss2Webhook/config"
	"GoRss2Webhook/core"
	"GoRss2Webhook/web/handler"
	"GoRss2Webhook/web/middleware"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log := logrus.WithField("source", "main")

	runPort, exist := os.LookupEnv(config.EnvRunPort)
	if !exist {
		runPort = "9090"
	}

	config.InitLog()
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	middleware.SetMiddleware(router)
	handler.Handle(router)

	core.Init()

	srv := &http.Server{
		Addr:    fmt.Sprintf(`:%s`, runPort),
		Handler: router,
	}

	go func() {
		log.Infof(`Server is running at :%s`, runPort)
		_ = srv.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infoln("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Infoln("Server exit!")
}
