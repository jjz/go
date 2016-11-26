package main

import (
	_ "beeapi/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.DEV == "dev" {
		beego.SetStaticPath("/swagger", "swagger")
	}
	beego.Run()
}
