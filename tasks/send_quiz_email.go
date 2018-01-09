package tasks

// https://myaccount.google.com/lesssecureapps
// https://medium.com/@dhanushgopinath/sending-html-emails-using-templates-in-golang-9e953ca32f3d

import (
	//"fmt"
	"github.com/astaxie/beego/toolbox"
	utils_mailer "beego-ripple/utils/mailer"
	"github.com/astaxie/beego"
)

const TASK_SEND_QUIZ_EMAIL  = "send_quiz_email"

func init() {
	cronSpec := beego.AppConfig.String("quiz_cron")
	// 0 0 8 10 * *
	first_task := toolbox.NewTask(TASK_SEND_QUIZ_EMAIL, cronSpec, func() error {

		beego.Info("Start Task - send_quiz_email");
		utils_mailer.SendQuizEmailToAll()
		return nil
	})

	toolbox.AddTask(TASK_SEND_QUIZ_EMAIL, first_task)
	toolbox.StartTask()
}






