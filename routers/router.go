package routers

import (
	"github.com/astaxie/beego"
	"mytools/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/randtool", &controllers.RandToolController{})
	beego.Router("/sql", &controllers.SqlController{})
}
