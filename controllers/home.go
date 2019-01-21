package controllers

import (
	"nkyWebServer/models"

	"github.com/astaxie/beego"
)

// MainController 主控制
type MainController struct {
	beego.Controller
}

// Get 方法
func (c *MainController) Get() {
	if !checkAccountCookie(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "home.html"
	nodes, err := models.GetAllNodes()
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Nodes"] = nodes
	nodeDatas, err := models.GetAllNodeData()
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["NodeDatas"] = nodeDatas
}
