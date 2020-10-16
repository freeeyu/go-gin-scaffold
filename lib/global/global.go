package global

import (
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	log "go_api/lib/logger"
	"os"
	"strconv"
)

//MakeData 数据
type MakeData map[string]interface{}

var conf *goconfig.ConfigFile

//User 当前用户信息
var User UserModel

//Json json格式化
func Json(message string, data interface{}) gin.H {
	if data == nil {
		data = []string{}
	}
	return gin.H{"message": message, "data": data}
}

func initConfig() {
	var err error
	conf, err = goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Ins.Logger().WithField("error", err.Error()).Fatalln("配置文件不存在,请将config.ini放入程序执行目录下。")
		os.Exit(0)
	}
}

//Config 配置读取
func Config(section string, key string) string {
	if conf == nil {
		initConfig()
	}
	sec, err := conf.GetSection(section)
	if err != nil {
		log.Ins.Logger().WithField("error", err.Error()).Fatalln("配置文件读取错误。")
		os.Exit(0)
	}
	return sec[key]
}

//GetRedis 从连接池中获取一个redis连接
func GetRedis() *redis.Client {
	maxActive, _ := strconv.Atoi(Config("redis", "maxActive"))
	db, _ := strconv.Atoi(Config("redis", "db"))
	return redis.NewClient(&redis.Options{
		Addr:     Config("redis", "address"),
		Password: "",
		DB:       db,
		PoolSize: maxActive,
	})
}
