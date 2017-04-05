package controllers

import (
	"github.com/astaxie/beego"
)

type SvnController struct {
	beego.Controller
	BaseController
}

func (this *SvnController) CheckLogin() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.Redirect(302, "/login")
	}
}
func (this *SvnController) Get() {
	this.CheckLogin()
	this.TplNames = "svnadd.tpl"
}

func (this *SvnController) Post() {
	this.CheckLogin()
	path, info := this.GetString("path"), this.GetString("info")
	if path == "" || info == "" {
		this.Ctx.WriteString("请输入path或info")
		return
	}

	tpl := "success.tpl"
	// this.Data["Msg"] = "添加"
	// if s == 0 {
	// 	tpl = "fail.tpl"
	// } else {
	// 	this.Data["Url"] = "/"
	// }
	this.TplNames = tpl
}
