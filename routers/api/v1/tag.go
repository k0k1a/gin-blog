package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/k0k1a/go-gin-example/models"
	"github.com/k0k1a/go-gin-example/pkg/e"
	"github.com/k0k1a/go-gin-example/pkg/setting"
	"github.com/k0k1a/go-gin-example/pkg/util"
	"net/http"
	"strconv"
)

func GetTags(c *gin.Context) {

	maps, data := make(map[string]interface{}), make(map[string]interface{})
	name := c.Query("name")
	if name != "" {
		maps["name"] = name
	}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
		maps["state"] = state
	}

	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
	name := c.Query("name")
	state, _ := strconv.Atoi(c.DefaultQuery("state", "0"))
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName("name") {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

func EditTag(c *gin.Context) {
}

func DeleteTag(c *gin.Context) {
}
