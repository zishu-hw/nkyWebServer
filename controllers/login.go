package controllers

import (
	"io"
	"nkyWebServer/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")
	autoLogin := c.Input().Get("autoLogin") == "on"
	beego.Info(uname, pwd, autoLogin)

	// 验证表单
	if models.CheckAccount(uname, pwd) {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("uname", uname, maxAge, "/")
		c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		io.WriteString(c.Ctx.ResponseWriter, "true")
		return
	} else {
		// c.EnableRender = false
		// c.Redirect("/login", 302)
		io.WriteString(c.Ctx.ResponseWriter, "账户密码错误")
		return
	}
	// 重定向
	// c.Redirect("/", 302)
}

func checkAccountCookie(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		beego.Error(err)
		return false
	}
	uname := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	o := orm.NewOrm()
	var dbpwd string
	err = o.Raw("select f_userPwd from tb_manager where f_userName=?", uname).QueryRow(&dbpwd)
	if err != nil {
		beego.Info(err.Error())
		return false
	}
	if pwd == dbpwd {
		return true
	}
	return false
}
