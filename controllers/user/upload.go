package user

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type UploadController struct {
	beego.Controller
}

func (u *UploadController) Get() {
	u.TplName = "user_model/upload.html"
}

func (u *UploadController) Post() {
	f, h, _ := u.GetFile("file")
	defer f.Close()

	tiem_unix := time.Now().Unix() //获取时间戳

	save_name := fmt.Sprintf("%d_%s", tiem_unix, h.Filename)
	fmt.Println(save_name)
	u.SaveToFile("file", "upload/"+save_name) //保存到指定路径
	u.Ctx.WriteString("上传成功")
}
