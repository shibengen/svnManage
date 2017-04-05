package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var DB orm.Ormer
var wwwroot, www string
var url, tpl, msg = "", "", ""

type BaseController struct {
	beego.Controller
}

/*
* 成功跳转
 */
func (this *BaseController) Success(msg string, url string) {
	data := make(map[string]interface{})
	data["msg"] = msg
	data["url"] = url
	this.Data["data"] = data
	this.TplNames = "success.tpl"
}

/*
* 失败跳转
 */
func (this *BaseController) Error(msg string) {
	data := make(map[string]interface{})
	data["msg"] = msg
	this.TplNames = "fail.tpl"

}

/*
* Ajax返回
*
 */
func (this *BaseController) AjaxReturn(status int, msg string, data interface{}) {
	json := make(map[string]interface{})
	json["status"] = status
	json["msg"] = msg
	json["data"] = data
	this.Data["json"] = json
	this.ServeJson()
	return
}

func init() {
	dbuser := beego.AppConfig.String("mysqluser")
	dbpassword := beego.AppConfig.String("mysqlpass")
	//dbhost := beego.AppConfig.String("mysqlurls")
	dbname := beego.AppConfig.String("mysqldb")
	wwwroot = beego.AppConfig.String("wwwroot")
	www = beego.AppConfig.String("www")
	orm.RegisterDataBase("default", "mysql", dbuser+":"+dbpassword+"@/"+dbname+"?charset=utf8")
	//orm.Debug = true
	DB = GetDb()
}

//DB
func GetDb() (o orm.Ormer) {
	o = orm.NewOrm()
	o.Using("default")
	return o
}
func (this *BaseController) Get_one(id string) Bat_info {
	var lists Bat_info
	DB.Raw("SELECT * FROM bat_info WHERE id = ?", id).QueryRow(&lists)
	return lists
}

//接收全体类型
type Alltype interface{}
