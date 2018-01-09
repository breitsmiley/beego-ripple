package mailer

import (
	"net/smtp"
	//"log"
	"fmt"
	"github.com/astaxie/beego"
	"bytes"
	"mime/quotedprintable"
	//"strings"
	"time"
	"beego-ripple/models"
	"html/template"
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
		//fmt.Printf("smtp error: %s", err)
		beego.Warn("smtp error", err)
		return
	}
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


//--------------
type NoticeTplData struct {
	AppSiteName string
	AppSiteSchema string
	AppSiteURL  string
	Time  string
	NoticeNum  int
}

func SendNoticeEmailToPartner(userEmail string) {

	appSiteName := beego.AppConfig.String("app_site_name")
	appSiteSchema := beego.AppConfig.String("app_site_schema")
	appSiteURL := beego.AppConfig.String("app_site_url")

	user1Email := beego.AppConfig.String("user1")
	user2Email := beego.AppConfig.String("user2")

	if userEmail != user1Email && userEmail != user2Email {
		beego.Warn("Wrong userEmail")
		return
	}

	var partnerEmail string

	if userEmail == user1Email {
		partnerEmail = user2Email
	} else {
		partnerEmail = user1Email
	}

	data := new(NoticeTplData)
	data.AppSiteName = appSiteName
	data.AppSiteSchema = appSiteSchema
	data.AppSiteURL = appSiteURL
	data.Time = time.Now().Local().Format("02.01.2006 15:04")
	data.NoticeNum = 1

	to := []string{partnerEmail}
	sendNoticeToEmail(to, data)
}

func renderEmailNoticeTpl(data interface{}) (tplText string, err error) {
	t, err := template.ParseFiles("views/email_notice.html")
	if err != nil {
		beego.Warn(err)
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		beego.Warn(err)
		return "", err
	}

	return buf.String(), nil
}

func sendNoticeToEmail(to []string, data *NoticeTplData) {
	sender := GetSender()

	subject := "Ripple.Quiz - Your partner have just given the answer - " + data.Time

	message, err := renderEmailNoticeTpl(data)
	if err != nil {
		beego.Warn(err)
		return
	}

	body := sender.WriteHTMLEmail(to, subject, message)

	sender.SendMail(to, body)

	beego.Info("Notice Mail send: ", data);
}


//---------------


type QuizTplData struct {
	AppSiteName string
	AppSiteSchema string
	AppSiteURL  string
	Time  string
	QuizURL  string
	NoticeNum  int
}


func SendQuizEmailToAll() {

	appSiteName := beego.AppConfig.String("app_site_name")
	appSiteSchema := beego.AppConfig.String("app_site_schema")
	appSiteURL := beego.AppConfig.String("app_site_url")

	quizList := models.ActivateQuizForUsers()

	if quizList == nil {
		beego.Info("Empty quiz list to mailing");
		return
	}

	data := new(QuizTplData)

	for i := 0; i < len(quizList) ; i++ {

		data.AppSiteName = appSiteName
		data.AppSiteSchema = appSiteSchema
		data.AppSiteURL = appSiteURL
		data.Time = time.Now().Local().Format("02.01.2006 15:04")

		data.NoticeNum = quizList[i].Step
		if  data.NoticeNum > 10  {
			data.NoticeNum = 1
		}

		data.QuizURL = appSiteURL + beego.URLFor(
			"MainController.Quiz",
			":id", quizList[i].Id,
			":slug", quizList[i].Slug)

		to := []string{quizList[i].User.Email}
		sendQuizToEmail(to, data)
	}

}

func renderEmailQuizTpl(data interface{}) (tplText string, err error) {
	t, err := template.ParseFiles("views/email_quiz.html")
	if err != nil {
		beego.Warn(err)
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		beego.Warn(err)
		return "", err
	}

	return buf.String(), nil
}

func sendQuizToEmail(to []string, data *QuizTplData) {
	sender := GetSender()

	subject := "Ripple.Quiz - It's time to make your choice - " + data.Time

	message, err := renderEmailQuizTpl(data)
	if err != nil {
		beego.Warn(err)
		return
	}

	body := sender.WriteHTMLEmail(to, subject, message)

	sender.SendMail(to, body)

	beego.Info("Quiz send: ", data);
}

