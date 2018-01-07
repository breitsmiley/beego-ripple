package tasks

// https://myaccount.google.com/lesssecureapps

import (
	"fmt"
	//"time"
	//"net/mail"

	"github.com/astaxie/beego/toolbox"
	"beego-ripple/utils"
	//"beego-ripple/models"
	//"log"
	//"net/smtp"
)

const TASK_SEND_QUIZ_EMAIL  = "send_quiz_email"

func init() {

	utils.MailSend("breitsmiley@gmail.com", "Subject GO mailer test 2018", "GO mailer test 2018 body")


	first_task := toolbox.NewTask(TASK_SEND_QUIZ_EMAIL, "0/5 * * * * *", func() error {
		// this task will run every 30 seconds

		fmt.Println("Task send_quiz_email")
		//
		////smtp.
		//
		//
		//// Set up authentication information.
		//auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")
		//
		//// Connect to the server, authenticate, set the sender and recipient,
		//// and send the email all in one step.
		//to := []string{"recipient@example.net"}
		//msg := []byte("To: recipient@example.net\r\n" +
		//	"Subject: discount Gophers!\r\n" +
		//	"\r\n" +
		//	"This is the email body.\r\n")
		//err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
		//if err != nil {
		//	log.Fatal(err)
		//}

		//models.InitDb()
		//campaigns, err := models.GetFinishedCampaigns()
		//if err != nil {
		//	fmt.Println("Could not load finished campaigns with error:", err)
		//	return err
		//}
		//
		//if len(campaigns) == 0 {
		//	fmt.Println("No campaigns finished yet. Exiting.")
		//	return nil
		//}
		//
		//for _, campaign := range campaigns {
		//	result, err := models.SendCampaignNotification(campaign)
		//	if err != nil {
		//		fmt.Printf("\nCampaign %d could not be notified!\n", campaign.Id)
		//	}
		//
		//	if result {
		//		fmt.Printf("\nCampaign %d was notified!\n", campaign.Id)
		//	}
		//}
		//
		//fmt.Printf("\nNotification task ran at: %s\n", time.Now())
		return nil
	})

	toolbox.AddTask(TASK_SEND_QUIZ_EMAIL, first_task)
	toolbox.StartTask()


	//defer toolbox.StopTask()
}