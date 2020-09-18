package main

import (
	D "go_api/lib/database"
	G "go_api/lib/global"
	"go_api/lib/response"
	"go_api/module/example"
	"go_api/module/user"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	D.Init()
	//静态文件
	r.StaticFS("/upload", http.Dir("./upload"))

	authGroup := r.Group("/v1")
	authGroup.Use(auth)

	authGroup.PUT("api/user", user.Put)

	// authGroup.DELETE("api/user", user.Delete)
	//不需要权限验证的
	noAuthGroup := r.Group("/v1")

	//登录
	noAuthGroup.GET("api/user", user.Get)
	//注册
	authGroup.POST("api/user", user.Post)
	//例子
	noAuthGroup.POST("api/example", example.Upload)
	noAuthGroup.POST("api/example/redis", example.Redis)

	r.Run(G.Config("server", "port"))
}

func auth(c *gin.Context) {
	token := c.GetHeader("token")
	if len(token) <= 0 {
		c.JSON(response.TokenInvalid.Code, G.Json(response.TokenInvalid.Message, nil))
		c.Abort()
		return
	}
	userID, err := D.DBT("user").Where("token", token).Value("id")
	if err != nil {
		c.JSON(response.UserInvalid.Code, G.Json(response.UserInvalid.Message, nil))
		log.Println(err.Error())
		c.Abort()
		return
	}
	if userID == nil || userID.(int64) <= 0 {
		c.JSON(response.UserInvalid.Code, G.Json(response.UserInvalid.Message, nil))
		c.Abort()
		return
	}
	err = D.DB().Table(&G.User).Where("id", userID).Select()
	if err != nil {
		c.JSON(response.UserInvalid.Code, G.Json(response.UserInvalid.Message, nil))
		log.Println(err.Error())
		c.Abort()
		return
	}
	if G.User.UID <= 0 {
		c.JSON(response.UserInvalid.Code, G.Json(response.UserInvalid.Message, nil))
		c.Abort()
		return
	}
}
