package main

import (
	"fmt"
	"github.com/k0k1a/go-gin-example/pkg/setting"
	"github.com/k0k1a/go-gin-example/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("start server failed:%v", err)
	}
}
