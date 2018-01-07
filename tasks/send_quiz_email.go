package tasks

import (
	"fmt"
	//"time"
	"github.com/astaxie/beego/toolbox"
	//"beego-ripple/models"
)

const TASK_SEND_QUIZ_EMAIL  = "send_quiz_email"

func init() {
	first_task := toolbox.NewTask(TASK_SEND_QUIZ_EMAIL, "0/5 * * * * *", func() error {
		// this task will run every 30 seconds

		fmt.Println("Task send_quiz_email")

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