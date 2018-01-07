package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beego-ripple/controllers:MainController"] = append(beego.GlobalControllerRouter["beego-ripple/controllers:MainController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego-ripple/controllers:MainController"] = append(beego.GlobalControllerRouter["beego-ripple/controllers:MainController"],
		beego.ControllerComments{
			Method: "InitDb",
			Router: `/init/hde15knQw`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["beego-ripple/controllers:MainController"] = append(beego.GlobalControllerRouter["beego-ripple/controllers:MainController"],
		beego.ControllerComments{
			Method: "Quiz",
			Router: `/quiz/:id/:slug`,
			AllowHTTPMethods: []string{"post","get"},
			MethodParams: param.Make(),
			Params: nil})

}
