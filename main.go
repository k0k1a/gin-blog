package main

import (
	"fmt"
	"github.com/k0k1a/go-gin-example/models"
	"github.com/k0k1a/go-gin-example/pkg/gredis"
	"github.com/k0k1a/go-gin-example/pkg/logging"
	"github.com/k0k1a/go-gin-example/pkg/setting"
	"github.com/k0k1a/go-gin-example/routers"
	"log"
	"net/http"
)

func init() {
	setting.SetUp()
	models.SetUp()
	logging.SetUp()
	gredis.SetUp()
}

func main() {
	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("start server failed:%v", err)
	}
}
