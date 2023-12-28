package cont

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type ValidController struct {
	beego.Controller
}

type MyUser struct {
	Id       int    `form:"_"`
	UserName string `form:"user_name" valid:"Required"`
	PassWord string `form:"password" valid:"Required"`
}

func (v *ValidController) Get() {
	v.TplName = "cont/valid.html"

}

func (v *ValidController) Post() {
	var user MyUser
	v.ParseForm(&user)
	fmt.Println(user)

	valid := validation.Validation{}

	//重写报错
	var MessageTmpls = map[string]string{
		"Required": "不能为空",
		"Min":      "最小值 %d",
		"Max":      "最大值 %d",
		"Range":    "介于 %d ～ %d",
	}
	validation.SetDefaultMessage(MessageTmpls)

	b, err := valid.Valid(&user)
	if err != nil {
		v.Ctx.WriteString("结构体tag出错")
	}
	if !b {
		for _, err_msg := range valid.Errors {
			fmt.Println(err_msg.Key)
			fmt.Println(err_msg.Message)
			v.Ctx.WriteString(err_msg.Key[:8] + err_msg.Message)
			return
		}
	}

	v.Ctx.WriteString("登录成功")
}
