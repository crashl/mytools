package routers

import (
	"beegodemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	//下面的方法是同一个文件里面处理不同的方法，自定义之后将不会访问默认的Get和Post方法
	beego.Router("/login", &controllers.RegisterController{},"get:ShowLogin;post:HandleLogin")
}
