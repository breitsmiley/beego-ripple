package routers

import (
	"beego-ripple/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{}, "get:Index")
    //beego.Router("/init", &controllers.MainController{}, "get:InitDb")
    //beego.Router("/quiz/:slug", &controllers.MainController{}, "get:Quiz")

	beego.Include(&controllers.MainController{})


}
