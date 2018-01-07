package routers

import (
	"github.com/astaxie/beego"
	"beego-ripple/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{}, "get:Index")
    //beego.Router("/init", &controllers.MainController{}, "get:InitDb")
    //beego.Router("/quiz/:slug", &controllers.MainController{}, "get:Quiz")

	beego.Include(&controllers.MainController{})


}
