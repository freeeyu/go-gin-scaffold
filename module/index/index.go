package index

import (
	"github.com/gin-gonic/gin"
	G "go_api/lib/global"
	"go_api/lib/response"
)

//Index 首页
func Index(c *gin.Context) {
	content := G.MakeData{
		"title": "Gin 模板测试",
		"desc":  "模板文件路径/templates/",
	}
	c.HTML(response.HTTPStatusOK, "index.tmpl", content)
}
