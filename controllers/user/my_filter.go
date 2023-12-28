package user

import "github.com/astaxie/beego/context"

func MyFilter(ctx *context.Context) {
	user_name := ctx.Input.Session("user_name")

	if user_name == nil {
		ctx.WriteString("未登录")
	}
}
