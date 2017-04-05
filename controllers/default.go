package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"io"
)

type MainController struct {
	beego.Controller
	BaseController
}

func (this *MainController) Get() {
	this.CheckLogin()
	v := this.GetSession("username")
	var list Alltype
	//var num int64
	list = this.getsvn()
	this.Data["Username"] = v
	this.Data["list"] = list
	this.TplNames = "index.tpl"
}

//getsvn
func (this *MainController) getsvn() Alltype {
	uid := this.GetSession("uid")
	var lists []Bat_info
	DB.Raw("SELECT * FROM bat_info WHERE uid = ?", uid).QueryRows(&lists)
	return lists
}

type Bat_info struct {
	Id   int
	Uid  int
	Path string
	Info string
}

//32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}

func (this *MainController) CheckLogin() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.Redirect(302, "/login")
	}
}
