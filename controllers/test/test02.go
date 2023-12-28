package test

import "github.com/astaxie/beego"

type Test02Controller struct {
	beego.Controller
}

func (t *Test02Controller) Get() {
	t.TplName = "test/test02.html"

	t.Data["x"] = "x"
	t.Data["y"] = ""
	t.Data["z"] = 19
	t.Data["arr"] = []int{0, 1, 2}
	t.Data["map_data"] = map[string]interface{}{"name": "DouSha", "age": 19}
	t.Data["is_ok"] = true
	t.Data["age"] = 18

}
