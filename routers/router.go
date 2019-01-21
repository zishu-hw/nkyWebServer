package routers

import (
	"nkyWebServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login/ajax", &controllers.LoginAjaxController{})
}
