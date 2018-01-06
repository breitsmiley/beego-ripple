package controllers

import (
	"github.com/astaxie/beego"
	"beego-ripple/models"
	//"api/models"
	//"fmt"
	"fmt"
)

type MainController struct {
	beego.Controller
}
// @router / [get]
func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// @router /init [get]
func (c *MainController) InitDb() {

	models.InitDb();

	c.Redirect(c.URLFor("MainController.Index"), 302)
}

// @router /quiz/:id/:slug [get]
func (c *MainController) Quiz() {

	//slug := c.Ctx.Input.Param(":slug")
	//id := c.Ctx.Input.Param(":id")

	id, err := c.GetInt(":id")
	if err != nil {
		fmt.Println(err)
		c.Redirect(c.URLFor("MainController.Index"), 302)
	}

	slug := c.GetString(":slug")


	fmt.Println(id)
	fmt.Println(slug)


	quiz , err := models.FindNewQuiz(id, slug);

	if err != nil {
		fmt.Println(err)
		//return quiz, err
		c.Redirect(c.URLFor("MainController.Index"), 302)
	}


	c.Data["Website"] = quiz.User.Email
	c.Data["Email"] = quiz.User.Email
	c.TplName = "index.tpl"
}


//// @router /quiz/:id/:slug [get]
//func (c *MainController) Quiz(id int, slug string) {
//
//	fmt.Println(id)
//	fmt.Println(slug)
//
//	//slug := c.Ctx.Input.Param(":slug")
//	//id := c.Ctx.Input.Param(":id")
//
//	quiz , err := models.FindNewQuiz(id, slug);
//
//	if err != nil {
//		fmt.Println(err)
//		//return quiz, err
//		c.Redirect(c.URLFor("MainController.Index"), 302)
//	}
//
//	//c.Redirect(c.URLFor("MainController.Index"), 302)
//
//	c.Data["Website"] = quiz.User.Email
//	c.Data["Email"] = quiz.User.Email
//	c.TplName = "index.tpl"
//
//
//	c.Render()
//}
