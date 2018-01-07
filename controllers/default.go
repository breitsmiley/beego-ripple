package controllers

import (
	"github.com/astaxie/beego"
	"beego-ripple/models"
	"fmt"
)

type MainController struct {
	beego.Controller
}

// @router / [get]
func (c *MainController) Index() {

	data, err := models.GetQuizViewData()

	if err != nil {
		fmt.Println(err)
		c.Abort("403")
	}
	c.Data["data"] = &data
	c.TplName = "index.html"

	failQuiz, _ := models.FindFailQuiz()
	if failQuiz != nil {
		fmt.Println(failQuiz.User)
		c.Data["fail"] = &failQuiz
		c.Data["user1"] = beego.AppConfig.String("user1")
		c.Data["user2"] = beego.AppConfig.String("user2")
		c.TplName = "fail.html"
	}

}

// @router /init/hde15knQw [get]
func (c *MainController) InitDb() {

	models.InitDb();

	c.Redirect(c.URLFor("MainController.Index"), 302)
}

// @router /quiz/:id/:slug [post,get]
func (c *MainController) Quiz() {

	// Common for GET and POST
	//-----------------------
	c.TplName = "quiz.tpl"

	//slug := c.Ctx.Input.Param(":slug")
	//id := c.Ctx.Input.Param(":id")

	id, err := c.GetInt(":id")
	if err != nil {
		fmt.Println(err)
		c.Redirect(c.URLFor("MainController.Index"), 302)
	}

	slug := c.GetString(":slug")


	failQuiz, _ := models.FindFailQuiz()
	if failQuiz != nil {
		c.Redirect(c.URLFor("MainController.Index"), 302)
	}

	quiz , err := models.FindActiveQuiz(id, slug);
	if err != nil {
		fmt.Println(err)
		c.Redirect(c.URLFor("MainController.Index"), 302)
	}


	// Common for GET and POST
	//-----------------------
	if (c.Ctx.Request.Method == "POST") {

		btn := c.GetString("btn")

		status := models.QUIZ_STATUS_OK
		if (btn == "no") {
			status = models.QUIZ_STATUS_FAIL
		}


		models.UpdateQuizStatus(quiz, status)

		c.Redirect(c.URLFor("MainController.Index"), 302)
	}
}