package main

import (
	"BeeGo/controllers/test"
	"BeeGo/controllers/user"
	_ "BeeGo/models/user_model"
	_ "BeeGo/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//读取配置文件
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	database := beego.AppConfig.String("database")
	//连接数据库
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", username, password, host, port, database)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", datasource)
	fmt.Println("连接数据库成功")

	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		panic(err)
	}
}
func main() {
	//自定义函数
	beego.AddFuncMap("substr_rep", test.SubStrAndRep)

	//过滤器
	beego.InsertFilter("/cms/*", beego.BeforeRouter, user.MyFilter)

	orm.RunCommand()

	//Logs
	logs.SetLogger(logs.AdapterConsole) //打印到控制台
	//logs.SetLogger(logs.AdapterFile, `{"filename":"logs/BeeGo.log"}`) //输出到文件
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/BeeGo.log","separate":["error","info"]}`) //输出到分级文件

	beego.Run()
}
