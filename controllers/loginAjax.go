package controllers

import (
	"io"

	"github.com/astaxie/beego"
)

type LoginAjaxController struct {
	beego.Controller
}

func (c *LoginAjaxController) Get() {
	io.WriteString(c.Ctx.ResponseWriter, "账户密码错误")
}
