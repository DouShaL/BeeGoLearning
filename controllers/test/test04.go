package test

import (
	"github.com/astaxie/beego"
)

type Test04Controller struct {
	beego.Controller
}

//自定义模板函数
func SubStrAndRep(str string, star_num, end_num int) string {
	bt := []rune(str)
	if end_num < len(bt) {
		ret_str := string(bt[star_num:end_num]) + "..."
		return ret_str
	} else {
		return str
	}
}

func (t *Test04Controller) Get() {
	t.TplName = "test/test04.html"

}
