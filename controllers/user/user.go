package user

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}
type UserInfo struct {
	User_name string `form:"user_name"`
	Password  string `form:"password"`
}

func (u *UserController) Get() {
	u.TplName = "user/user.html"
	//id := u.GetString(":id")	//router内有则加：
	//id, err := u.GetInt("id")
	//if err != nil {
	//	u.Ctx.WriteString("ID为int.")
	//}
	//u.Data["id"] = id
}

func (u *UserController) Post() {
	user := UserInfo{}
	u.ParseForm(&user)
	fmt.Println(user.User_name)
	fmt.Println(user.Password)

	u.SetSession("user_name", "DouSha")
	ret := u.GetSession("user_name")
	fmt.Println("session已设置:", ret)
	//u.DelSession("user_name")
	//ret = u.GetSession("user_name")
	//fmt.Println(ret)

	u.Ctx.WriteString("登录成功")
	//ret := map[string]interface{}{"code": 200, "msg": "登录成功"}
	//u.Data["json"] = ret
	//u.ServeJSON()
}
