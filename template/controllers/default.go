package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Jade() {
	c.Data["content"] = "this is jade template"
	c.TplName = "home.jade"
}
func (c *MainController)Ace() {
	c.Data["content"] = "ace"
	c.TplName = "home.ace"
}
func (c *MainController)Index() {

	c.TplName = "index.html"
}
