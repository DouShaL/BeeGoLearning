package test

import "github.com/astaxie/beego"

type Test01Controller struct {
	beego.Controller
}

func (t *Test01Controller) Get() {
	t.TplName = "test/test01.html"

	t.Data["name"] = "DouSha"
	t.Data["arr"] = []int{0, 1, 2, 3}
	t.Data["arr1"] = []int{1}

}
