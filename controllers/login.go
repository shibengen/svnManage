package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
	BaseController
}

func (this *LoginController) Get() {
	v := this.GetSession("username")
	if v != nil {
		this.Ctx.Redirect(302, "/")
	}
	this.TplNames = "login.tpl"
}
func (this *LoginController) Post() {
	u := user{this.GetString("username"), this.GetString("password")}
	if u.username == "" || u.password == "" {
		this.Ctx.WriteString("用户名或密码不能为空")
		return
	}
	s := u.Login()
	if s > 0 {
		this.SetSession("username", u.username)
		this.SetSession("uid", s)
		this.Ctx.Redirect(302, "/")
		msg = "登录成功"
		url = "/"
		tpl = "success.tpl"
	} else {
		msg = "登录失败"
		tpl = "fail.tpl"
	}
	this.Data["msg"] = msg
	this.Data["url"] = url
	this.TplNames = tpl

}

type user struct {
	username string
	password string
}

func (u *user) Login() int {
	var userInfo User
	err := DB.Raw("SELECT uid,username FROM userinfo WHERE username=? AND password=?", u.username, GetMd5String(u.password)).QueryRow(&userInfo)
	if err != nil {
		return 0
	}
	return userInfo.Uid
}

type User struct {
	Uid      int `orm:"auto"`
	UserName string
}
