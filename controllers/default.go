package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type User struct {
	Id   int
	Name string
	Age  int
}

func (c *MainController) Get() {
	//c.TplName = "index.tpl"
	c.TplName = "index.html"

	//渲染字符串到页面,不加载模板
	//c.Ctx.WriteString("Hello,BeeGo!")
	c.Data["Website"] = "DouSha"
	c.Data["Email"] = "2683382910@qq.com"

	//var user_model User
	//user_model.Id = 1
	//user_model.Name = "DouSha"
	//user_model.Age = 19
	user := User{
		Id:   1,
		Name: "DouSha",
		Age:  19,
	}
	c.Data["user_model"] = user

	arr := [...]int{0, 1, 2, 3, 4, 5}
	c.Data["arr"] = arr

	users := [3]User{
		{Id: 1, Name: "  Bee", Age: 18},
		{Id: 2, Name: "   Go", Age: 19},
		{Id: 3, Name: "BeeGo", Age: 20},
	}
	c.Data["users"] = users

	//map_data := map[string]string{
	//	"name": "DouSha",
	//	"QQ": "2683382910",
	//}
	map_data := map[string]interface{}{
		"name": "DouSha",
		"QQ":   2683382910,
	}
	c.Data["map_data"] = map_data

	map_struct := map[string]User{
		"user1": {
			Id:   1,
			Name: "Block",
			Age:  1,
		},
		"user2": {
			Id:   2,
			Name: "Chain",
			Age:  2,
		},
	}
	c.Data["map_struct"] = map_struct
}

//func (c *MainController) Post()  {
//
//}
