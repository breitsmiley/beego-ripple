package utils

import (
	"net/smtp"
	//"log"
	"fmt"
	"github.com/astaxie/beego"
	"bytes"
	"mime/quotedprintable"
	//"strings"
)


type Sender struct {
	User     string
	Password string
	SmtpHost string
	SmtpPort string
}


func GetSender() Sender {

	return Sender{
	beego.AppConfig.String("smtp_user"),
beego.AppConfig.String("smtp_pass"),
beego.AppConfig.String("smtp_host"),
beego.AppConfig.String("smtp_port"),
	}
}


func (sender Sender) SendMail(to []string, body string) {

	//msg := "From: " + sender.User + "\n" +
	//	"To: " + strings.Join(to, ",") + "\n" +
	//	"Subject: " + subject + "\n\n" +
	//	body

	err := smtp.SendMail(sender.SmtpHost+":"+sender.SmtpPort,
		smtp.PlainAuth("", sender.User, sender.Password, sender.SmtpHost),
		sender.User, to, []byte(body))

	if err != nil {
		//log.Printf("smtp error: %s", err)
		fmt.Printf("smtp error: %s", err)
		return
	}

	// todo
	//log.Print("sent, visit http://foobarbazz.mailinator.com")
}


func (sender Sender) WriteEmail(dest []string, contentType, subject, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = sender.User

	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}

	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=UTF-8", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func (sender *Sender) WriteHTMLEmail(dest []string, subject, bodyMessage string) string {

	return sender.WriteEmail(dest, "text/html", subject, bodyMessage)
}

func (sender *Sender) WritePlainEmail(dest []string, subject, bodyMessage string) string {

	return sender.WriteEmail(dest, "text/plain", subject, bodyMessage)
}
