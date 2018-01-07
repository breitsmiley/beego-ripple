package utils

import (
	"net/smtp"
	//"log"
	"fmt"
	"github.com/astaxie/beego"
)

func MailSend(to string, sbj string, body string) {
	from := beego.AppConfig.String("mailer_from")
	pass := beego.AppConfig.String("mailer_pass")
	smtp_host := beego.AppConfig.String("mailer_smtp_host")
	smtp_port := beego.AppConfig.String("mailer_smtp_port")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + sbj + "\n\n" +
		body

	err := smtp.SendMail(smtp_host+":"+smtp_port,
		smtp.PlainAuth("", from, pass, smtp_host),
		from, []string{to}, []byte(msg))

	if err != nil {
		//log.Printf("smtp error: %s", err)
		fmt.Printf("smtp error: %s", err)
		return
	}

	// todo
	//log.Print("sent, visit http://foobarbazz.mailinator.com")
}
