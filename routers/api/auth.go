package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/k0k1a/go-gin-example/models"
	"github.com/k0k1a/go-gin-example/pkg/e"
	"github.com/k0k1a/go-gin-example/pkg/logging"
	"github.com/k0k1a/go-gin-example/pkg/util"
	"net/http"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok, _ := valid.Valid(&auth{username, password}); ok {
		if isExist := models.CheckAuth(username, password); isExist {
			if token, err := util.GenerateToken(username, password); err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			//log.Println(err.Key, err.Message)
			logging.Info(err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
