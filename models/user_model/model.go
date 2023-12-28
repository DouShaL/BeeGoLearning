package user_model

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int    `orm:"pk;auto"`      //主键；自增
	Name       string `orm:"index;unique"` //索引
	Age        int
	Addr       string     `orm:"null;column(address)"`            //允许为null
	Desc       string     `orm:"size(2000)"`                      //string长度2000
	Price      float64    `orm:"digits(12);decimals(4)"`          //长度12；小数点4
	CreateTime time.Time  `orm:"auto_now_add;type(datetime)"`     //时间戳
	Status     int        `orm:"default(1);description(1启用；0停用)"` //设置默认值1;设置注释
	XXX        string     `orm:"-"`                               //忽略
	Profile    *Profile   `orm:"rel(one)"`                        //一对一正向有外键
	Articles   []*Article `orm:"reverse(many)"`
}
type Profile struct {
	Id     int
	IdCard string
	Cover  string
	User   *User `orm:"reverse(one)"` //反向无外键
}
type Article struct {
	Id      int
	Title   string
	Content string `orm:"size(2000)"`
	User    *User  `orm:"rel(fk)"` //一对多
}
type Post struct {
	Id    int
	Title string
	Tags  []*Tag `orm:"rel(m2m)"` //多对多
}
type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func (u *User) TableName() string {
	return "sys_user"
}
func (p *Profile) TableName() string {
	return "sys_profile"
}
func (a *Article) TableName() string {
	return "sys_article"
}
func (p *Post) TableName() string {
	return "sys_post"
}
func (t *Tag) TableName() string {
	return "sys_tag"
}

func init() {
	orm.RegisterModel(new(User), new(Profile), new(Article), new(Post), new(Tag))
}
