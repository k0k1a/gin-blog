package app

import (
	"github.com/gin-gonic/gin"
	"github.com/k0k1a/go-gin-example/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errorCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errorCode,
		"msg":  e.GetMsg(errorCode),
		"data": data,
	})
	return
}
