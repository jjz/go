package routers

import (
	"github.com/astaxie/beego"
	"Go/template/controllers"
)

func init() {
	beego.Router("/jade", &controllers.MainController{}, "Get:Jade")
	beego.Router("/ace", &controllers.MainController{}, "Get:Ace")
	beego.Router("/", &controllers.MainController{}, "Get:Index")
}
