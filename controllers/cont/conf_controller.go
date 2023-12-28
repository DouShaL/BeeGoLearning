package cont

import "github.com/astaxie/beego"

type ConfController struct {
	beego.Controller
}

func (c *ConfController) Get() {
	c.TplName = "cont/conf_controller.html"

	httpport, _ := beego.AppConfig.Int("httpport")
	c.Data["httpport"] = httpport
	username := beego.AppConfig.String("username")
	c.Data["username"] = username
	password := beego.AppConfig.String("password")
	c.Data["password"] = password
}
