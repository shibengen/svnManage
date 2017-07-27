package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"os/exec"
	"time"
	//"strings"
)

type ExecController struct {
	beego.Controller
	BaseController
}

func (this *ExecController) CheckLogin() {
	v := this.GetSession("username")
	if v == nil {
		this.Ctx.Redirect(302, "/login")
	}
}
func (this *ExecController) Get() {
	this.CheckLogin()
	id := this.GetString("id")
	if id == "" {
		this.Error("参数错误")
	}
	var info Bat_info
	var num int64
	info = this.Get_one(id)
	userFile := info.Path + "_id_" + id + ".bat"
	out, err := exec.Command(userFile).Output()
	str := string(out)
	if err != nil {
		fmt.Println(err)
		msg = "更新失败" + err.Error()
		tpl = "fail.tpl"
	} else {
		msg = "更新成功"
		url = "/"
		tpl = "exec.tpl"
		this.Data["message"] = userFile
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	l := len(str)

	if l > 0 {
		f := "./log/" + id + ".log"
		fout, err := os.Create(f)
		if err != nil {
			fmt.Println(userFile, err)
			return
		}
		defer fout.Close()
		str = "[ " + t + " ]已更新以下文件:" + str[108:]
		fout.WriteString(str)
	} else {
		str = "没有文件被更新"
	}

	this.Data["msg"] = msg
	this.Data["str"] = str
	this.Data["url"] = url
	this.TplNames = tpl
}
