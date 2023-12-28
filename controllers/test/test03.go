package test

import (
	"github.com/astaxie/beego"
	"time"
)

type Test03Controller struct {
	beego.Controller
}

func (t *Test03Controller) Get() {
	t.TplName = "test/test03.html"

	t.Data["time_data"] = time.Now()
	t.Data["map_data"] = map[string]interface{}{"name": "DouSha", "age": 19}
}
