package user

import (
	"github.com/astaxie/beego"
)

type OtherTypeController struct {
	beego.Controller
}
type Food struct {
	Id      int
	Name    string
	Quailty int
}

func (o *OtherTypeController) Get() {
	o.TplName = "user_model/othertype.html"

	food := Food{Id: 1, Name: "Apple", Quailty: 1}

	//json格式
	o.Data["json"] = food
	o.ServeJSON()
	//xml格式
	//o.Data["xml"] = food
	//o.ServeXML()
	//yaml格式
	//o.Data["yaml"] = food
	//o.ServeYAML()

}
