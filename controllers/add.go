package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type AddController struct {
	beego.Controller
	BaseController
}

func (this *AddController) CheckLogin() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.Redirect(302, "/login")
	}
}
func (this *AddController) Get() {
	this.CheckLogin()
	id := this.GetString("id")
	var info Bat_info
	info = this.Get_one(id)
	this.Data["info"] = info
	this.TplNames = "add.tpl"
}
func (this *AddController) Post() {
	this.CheckLogin()
	path := this.GetString("path")
	info := this.GetString("info")
	num := this.Save(path, info)
	url, tpl, msg := "", "", ""

	if num > 0 {
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
func (this *AddController) Save(path string, info string) int64 {
	uid := this.GetSession("uid")
	res, err := DB.Raw("INSERT INTO  bat_info (uid,path,info) VALUES (?,?,?)", uid, path, info).Exec()
	if err != nil {
		fmt.Println(err)
	}
	num, _ := res.RowsAffected()
	return num
}
