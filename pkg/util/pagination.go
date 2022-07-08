package util

import (
	"github.com/gin-gonic/gin"
	"github.com/k0k1a/go-gin-example/pkg/setting"
	"strconv"
)

// GetPage 得到分页
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}
