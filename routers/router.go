package routers

import (
	"BeeGo/controllers"
	"BeeGo/controllers/cont"
	"BeeGo/controllers/logs"
	"BeeGo/controllers/orm"
	"BeeGo/controllers/test"
	"BeeGo/controllers/user"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &controllers.MainController{})
	//beego.Router("/user_model/?:id:int", &controllers.UserController{})
	beego.Router("/upload", &user.UploadController{})
	beego.Router("/othertype", &user.OtherTypeController{})
	beego.Router("/conf", &cont.ConfController{})
	beego.Router("/valid", &cont.ValidController{})
	beego.Router("/exper", &orm.ExperController{})
	beego.Router("/query_table", &orm.QueryTableController{})
	beego.Router("/logs", &logs.LogsController{})

	//自定义路由
	beego.Router("/router", &cont.RouterController{}, "Get:List")

	//过滤器
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &user.UserController{})
	beego.Router("/add_user", &user.UserController{})

	//需要登录的
	beego.Router("/cms/test01", &test.Test01Controller{})
	beego.Router("/cms/test02", &test.Test02Controller{})
	beego.Router("/cms/test03", &test.Test03Controller{})
	beego.Router("/cms/test04", &test.Test04Controller{})

}
