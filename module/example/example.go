package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	G "go_api/lib/global"
	"go_api/lib/response"
	"go_api/lib/tools"
	V "go_api/lib/valid"
	"os"
	"time"
)

var valid = V.Validate{}

//Upload 上传文件示例
func Upload(c *gin.Context) {
	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(response.HTTPStatusFaild, G.Json(err.Error(), nil))
		return
	}
	if message, ok := valid.Check("上传图片", image, []string{"image"}); !ok {
		c.JSON(response.HTTPStatusFaild, G.Json(message, nil))
		return
	}
	dir, _ := os.Getwd()
	//目录不存在需要生成
	md5 := tools.MD5(time.Now().UnixNano())
	err = c.SaveUploadedFile(image, fmt.Sprintf("%s/upload/%s.jpg", dir, md5))
	if err != nil {
		c.JSON(response.HTTPStatusFaild, G.Json(err.Error(), nil))
		return
	}
	baseURL := fmt.Sprintf("%s%s%s", G.Config("server", "url"), G.Config("server", "port"), fmt.Sprintf("/upload/%s.jpg", md5))
	c.JSON(response.HTTPStatusOK, G.Json("数据更新成功", G.MakeData{"image": baseURL}))
}

//Redis redis使用示例
func Redis(c *gin.Context) {
	input := c.PostForm("input")
	redis := G.GetRedis()
	expire := 3600 * time.Second
	redis.Set("input", input, expire)
}
