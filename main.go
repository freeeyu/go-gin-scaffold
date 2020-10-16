package main

import (
	"github.com/gin-gonic/gin"
	D "go_api/lib/database"
	G "go_api/lib/global"
	log "go_api/lib/logger"
	"go_api/lib/response"
	"go_api/module/example"
	"go_api/module/index"
	"go_api/module/user"
	"net/http"
)

func init() {
	//init database
	// D.Init()
	//init log
	// log.Ins.Init()
	gin.DisableConsoleColor()
}

func main() {
	r := gin.New()
	r.Use(log.Ins.Gin())
	log.Ins.Logger().WithField("记录到日志的变量", r.BasePath()).Warningln("警告日志测试")
	log.Ins.Logger().WithFields(log.Ins.Fields(G.MakeData{"test": "yes", "key": "value"})).Errorln("错误日志测试-fields")
	//静态文件
	r.StaticFS("/statics/", http.Dir("./statics/"))
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

	//页面
	//加载模板,如果是/templates/User/index.tmpl的路径格式,则改成templates/**/*.tmpl
	r.LoadHTMLGlob("templates/*.tmpl")
	r.GET("/", index.Index)
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
		log.Ins.Logger().WithField("error", err.Error()).Errorln(err.Error())
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
		log.Ins.Logger().WithField("error", err.Error()).Errorln(err.Error())
		c.Abort()
		return
	}
	if G.User.UID <= 0 {
		c.JSON(response.UserInvalid.Code, G.Json(response.UserInvalid.Message, nil))
		c.Abort()
		return
	}
}
