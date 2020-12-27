package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "seckill-backend/routers"
)

func init() {
	beego.SetStaticPath("/img","static/img")
	beego.SetStaticPath("/css","static/css")
	beego.SetStaticPath("/js","static/js")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	ds :=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		beego.AppConfig.String("db::user"),
		beego.AppConfig.String("db::password"),
		beego.AppConfig.String("db::host"),
		beego.AppConfig.String("db::port"),
		beego.AppConfig.String("db::database"))
	dbConnectErr := orm.RegisterDataBase("default", "mysql", ds)
	if dbConnectErr != nil {
		panic(dbConnectErr)
	}
	maxIdle, _ := beego.AppConfig.Int("db::maxIdle")
	maxConn, _ := beego.AppConfig.Int("db::maxConn")
	orm.SetMaxIdleConns("default", maxIdle)
	orm.SetMaxOpenConns("default", maxConn)
	// 自动建表
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

func main() {
	beego.Run()
}

