package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	this.DelSession("username")
	this.DelSession("uid")
	this.Ctx.Redirect(302, "/login")

}
