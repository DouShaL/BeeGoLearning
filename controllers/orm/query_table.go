package orm

import (
	"BeeGo/models/user_model"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type QueryTableController struct {
	beego.Controller
}

func (q *QueryTableController) Get() {
	q.TplName = "orm/query_table.html"
	o := orm.NewOrm()
	user := user_model.User{}
	users := []user_model.User{}
	users1 := []user_model.User{}
	qs := o.QueryTable(new(user_model.User))
	//One:一条数据
	//qs.One(&user)
	//All:全部数据
	qs.All(&users)
	//Filter:查询条件
	qs.Filter("name", "BeeGo").One(&user)
	//Exclude:排除条件
	qs.Exclude("name", "BeeGo").One(&user)
	//Limit(x,y):从第y条数据查询x条
	qs.Limit(2, 1).All(&users)
	//Offset(x):从第x条开始查询
	qs.Offset(1).All(&users)
	//GroupBy:分组（只查找一条）
	qs.GroupBy("name").All(&users1)
	//OrderBy:排序(默认升序 -降序)
	qs.OrderBy("-age").All(&users1)
	//Distint:去重(在保存地址后制定字段)
	qs.Distinct().All(&users1, "address")
	//Count:统计数据数量
	count, _ := qs.Exclude("age", 2).Count()
	//Exist:查找是否存在
	exist := qs.Filter("age", 19).Exist()
	//Updata:更新数据
	qs.Filter("age", 2).Update(orm.Params{
		"price": 222,
	})
	//Delete:删除数据
	num, _ := qs.Filter("age", 666).Delete()
	fmt.Printf("删除了%d条数据\n", num)
	//Insert:插入数据
	//insert_profile := user.Profile
	//insert_profile.Id = 4
	//insert_profile.User = nil
	//insert_user := user_model.User{
	//	Id:      4,
	//	Name:    "Lan",
	//	Profile: insert_profile,
	//}
	//o.Insert(&insert_user)

	q.Data["user"] = user
	q.Data["users"] = users
	q.Data["users1"] = users1
	q.Data["count"] = count
	q.Data["exist"] = exist
}
