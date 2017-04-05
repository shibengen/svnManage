package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
)

type EditController struct {
	beego.Controller
	BaseController
}

func (this *EditController) CheckLogin() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.Redirect(302, "/login")
	}
}
func (this *EditController) Get() {
	this.CheckLogin()
	id := this.GetString("id")
	var info Bat_info
	info = this.Get_one(id)
	this.Data["info"] = info
	this.TplNames = "edit.tpl"
}
func (this *EditController) Post() {
	this.CheckLogin()
	id := this.GetString("id")
	path := this.GetString("path")
	info := this.GetString("info")
	num := this.Update(path, info, id)

	if num > 0 {
		//修改BAT文件
		userFile := path + "_id_" + id + ".bat"
		fout, err := os.Create(userFile)
		if err != nil {
			fmt.Println(userFile, err)
			return
		}
		defer fout.Close()
		fout.WriteString("svn update " + wwwroot + path + "\\" + www)
		msg = "成功"
		url = "/"
		tpl = "success.tpl"
	} else {
		msg = "失败"
		tpl = "fail.tpl"
	}
	this.Data["msg"] = msg
	this.Data["url"] = url
	this.TplNames = tpl
	//var info Bat_info

}
func (this *EditController) Update(path string, info string, id string) int64 {
	res, err := DB.Raw("UPDATE bat_info SET path = ? , info=? where id=?", path, info, id).Exec()
	if err != nil {
		fmt.Println(err)
	}
	num, _ := res.RowsAffected()
	return num
}
