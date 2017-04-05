package main

import (
	"github.com/astaxie/beego"
	_ "svnmanage/controllers"
	_ "svnmanage/routers"
)

func main() {
	beego.Run()
}
