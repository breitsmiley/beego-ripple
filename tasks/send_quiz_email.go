package tasks

// https://myaccount.google.com/lesssecureapps
// https://medium.com/@dhanushgopinath/sending-html-emails-using-templates-in-golang-9e953ca32f3d

import (
	"fmt"
	"html/template"
	"github.com/astaxie/beego/toolbox"
	"beego-ripple/utils"
	"bytes"
	"time"
	"github.com/astaxie/beego"
	"beego-ripple/models"
)

const TASK_SEND_QUIZ_EMAIL  = "send_quiz_email"


type QuizTplData struct {
	AppSiteName string
	AppSiteSchema string
	AppSiteURL  string
	Time  string
	QuizURL  string
}

func init() {
	first_task := toolbox.NewTask(TASK_SEND_QUIZ_EMAIL, "0 0 5 */1 * *", func() error {
		sendQuizEmailToAll()
		return nil
	})

	toolbox.AddTask(TASK_SEND_QUIZ_EMAIL, first_task)
	toolbox.StartTask()

}

func sendQuizEmailToAll() {

	fmt.Println("Task send_quiz_email")

	appSiteName := beego.AppConfig.String("app_site_name")
	appSiteSchema := beego.AppConfig.String("app_site_schema")
	appSiteURL := beego.AppConfig.String("app_site_url")

	quizList := models.ActivateQuizForUsers()

	if quizList == nil {
		fmt.Println("Empty quiz list to mailing")
		return
	}

	for i := 0; i < len(quizList) ; i++ {

		data := new(QuizTplData)
		data.AppSiteName = appSiteName
		data.AppSiteSchema = appSiteSchema
		data.AppSiteURL = appSiteURL
		data.Time = time.Now().Local().Format("02.01.2006 15:04")
		data.QuizURL = appSiteURL + beego.URLFor(
			"MainController.Quiz",
			":id", quizList[i].Id,
			":slug", quizList[i].Slug)

		fmt.Println(data)

		to := []string{quizList[i].User.Email}
		sendQuizToEmail(to, data)
	}

}

func renderEmailQuizTpl(data interface{}) (tplText string, err error) {
	t, err := template.ParseFiles("views/email_quiz.html")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
		return "", err
	}

	return buf.String(), nil
}

func sendQuizToEmail(to []string, data *QuizTplData) {
	sender := utils.GetSender()

	subject := "Ripple.Quiz - It's time to make your choice - " + data.Time

	message, err := renderEmailQuizTpl(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := sender.WriteHTMLEmail(to, subject, message)

	sender.SendMail(to, body)
}





