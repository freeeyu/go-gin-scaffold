package database

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
	log "go_api/lib/logger"
	// "github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	G "go_api/lib/global"
	"os"
)

// var db *sql.DB
var engin *gorose.Engin

func init() {
	var err error
	conn := fmt.Sprintf("%s:%s@%s?parseTime=true", G.Config("mysql", "username"), G.Config("mysql", "password"), G.Config("mysql", "host"))
	engin, err = gorose.Open(&gorose.Config{Driver: "mysql", Dsn: conn})
	if err != nil {
		log.Ins.Logger().WithField("error", err.Error()).Fatalln("数据库初始化失败")
		os.Exit(0)
	}
}

//DB 获取数据库连接
func DB() gorose.IOrm {
	if engin == nil {
		log.Ins.Logger().Fatalln("数据库初始化失败")
		os.Exit(0)
	}
	return engin.NewOrm()
}

//DBT 带表名的数据库连接
func DBT(table string) gorose.IOrm {
	if engin == nil {
		log.Ins.Logger().Fatalln("数据库初始化失败")
		os.Exit(0)
	}
	e := engin.NewOrm()
	if len(table) > 0 {
		e.Table(table)
	}
	return e
}
