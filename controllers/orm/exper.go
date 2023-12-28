package orm

import (
	"BeeGo/models/user_model"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ExperController struct {
	beego.Controller
}

func (e *ExperController) Get() {
	o := orm.NewOrm()
	//o.QueryTable("sys_user")
	qs := o.QueryTable(new(user_model.User))
	user := user_model.User{}
	//qs.Filter("name", "DouSha").One(&user)		  //默认等于查找：__exact
	//qs.Filter("name__contains", "Sha").One(&user) //包含查找：__contains
	//gt(大于) gte(大于等于) lt(小于) lte(小于等于)
	//qs.Filter("age__gt", 18).One(&user) //大于查找：__gt
	//qs.Filter("name__iendswith", "sha").One(&user) //结尾查找：__endswith(大小写敏感) __iendswith(大小写不敏感)
	//qs.Filter("name__istartswith", "dou").One(&user) //起始查找：__startswith(大小写敏感) __istartswith(大小写不敏感)
	//qs.Filter("age__in", 18, 19, 20).One(&user) //范围查找：__in
	qs.Filter("address__isnull", false).One(&user) //null查找：__isnull ture为空的 false不为空的

	fmt.Println(user)
	e.Data["user"] = user
	e.TplName = "orm/exper.html"
}
