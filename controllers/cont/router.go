package cont

import "github.com/astaxie/beego"

type RouterController struct {
	beego.Controller
}

func (r *RouterController) List() {
	r.TplName = "cont/router.html"

}
