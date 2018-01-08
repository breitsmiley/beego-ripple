package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
	"beego-ripple/utils"
)

const QUIZ_STATUS_NEW  = 0
const QUIZ_STATUS_ACTIVE  = 1
const QUIZ_STATUS_OK  = 2
const QUIZ_STATUS_FAIL  = 3

type User struct {
	Id          int `orm:"auto"`
	Email        string
}

type Quiz struct {
	Id    int    `orm:"auto"`
	User  *User  `orm:"rel(fk)"`
	Status int `orm:"default(0)"`
	Slug string
	Step int `orm:"default(0)"`
}

func init() {
	orm.RegisterModel(new(User), new(Quiz))
}

func InitDb() {

	beego.Info("Init DB")

	o := orm.NewOrm()
	//cnt, err := o.QueryTable("user").Count() // SELECT COUNT(*) FROM USER

	runMod := beego.AppConfig.String("runmode")
	if runMod != "dev" {
		beego.Info("Init DB Prohibited - Not DEV run mod")
		return
	}

	// Error.
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		beego.Warn(err)
		return
	}

	user1 := new(User)
	user1.Email = beego.AppConfig.String("user1")
	o.Insert(user1)
	fillQuizForUser(user1);


	user2 := new(User)
	user2.Email = beego.AppConfig.String("user2")
	o.Insert(user2)
	fillQuizForUser(user2);

	return
}

func FindActiveQuiz(id int, slug string) (quiz *Quiz, err error) {
	quiz = new(Quiz)
	o := orm.NewOrm()

	qs := o.QueryTable("quiz").
		Filter("id", id).
		Filter("slug", slug).
		Filter("status", QUIZ_STATUS_ACTIVE).
		Limit(1).
		RelatedSel()

	err = qs.One(quiz)

	if err != nil {
		beego.Warn(err)
		return nil, err
	}

	return quiz, nil
}

func FindFailQuiz() (quiz *Quiz, err error) {

	quiz = new(Quiz)
	o := orm.NewOrm()

	qs := o.QueryTable("quiz").
		Filter("status", QUIZ_STATUS_FAIL).
		Limit(1).
		RelatedSel()

	err = qs.One(quiz)

	if err != nil {
		beego.Warn(err)
		return nil, err
	}

	return quiz, nil
}

func fillQuizForUser(user * User)  {
	o := orm.NewOrm()

	for i := 0; i < 10; i++ {

		quiz := new(Quiz)
		quiz.User = user
		quiz.Status = QUIZ_STATUS_NEW
		quiz.Slug = utils.RandomString(10)
		//quiz.Step = 1 + 2 * i
		quiz.Step = i+1

		o.Insert(quiz)
	}
}


func UpdateQuizStatus(quiz * Quiz, status int)  {

	o := orm.NewOrm()

	if (status == QUIZ_STATUS_FAIL) {
		quiz.Status = status
		if num, err := o.Update(quiz); err == nil {
			beego.Info("Updated rows = ", num)
		}
	}

	if (status == QUIZ_STATUS_OK) {

		num, err := o.QueryTable("quiz").
			Filter("status", QUIZ_STATUS_ACTIVE).
			Filter("user", quiz.User.Id).
			Update(orm.Params{"status": QUIZ_STATUS_OK})
		beego.Info("Updated rows = ", num, err)
	}




}


func getProgressBarData(user User) (quizList []*Quiz, err error) {

	o := orm.NewOrm()

	qs := o.QueryTable("quiz").Filter("user", user).RelatedSel()
	num , err := qs.All(&quizList)

	if err != nil {
		beego.Warn(err)
		return nil, err
	}

	if num == 0 {
		return nil, err
	}

	return quizList, nil
}

type ProgressBarData struct {
	Percent int
}

type QuizViewData struct {
	Pb1 *ProgressBarData
	Pb2 *ProgressBarData
}

func prepareProgressBarData(quizList []*Quiz) (data *ProgressBarData) {

	okSteps := 0
	for i := range quizList {
		if (quizList[i].Status == QUIZ_STATUS_OK) {
			okSteps++
		}
	}

	data = new(ProgressBarData)
	data.Percent = okSteps * 100 / len(quizList)

	return data

}

func activateQuizForUser (user User) (quiz *Quiz, err error) {
	o := orm.NewOrm()

	quiz = new(Quiz)

	qs := o.QueryTable("quiz").
		Filter("user", user).
		Filter("status", QUIZ_STATUS_NEW).
		RelatedSel()

	err = qs.One(quiz)

	if err != nil {
		beego.Warn(err)
		return nil, err
	}

	quiz.Status = QUIZ_STATUS_ACTIVE
	if num, err := o.Update(quiz); err == nil {
		beego.Info("Updated rows = ", num, err)
	}

	return quiz, nil

}

func ActivateQuizForUsers() (quizList []*Quiz) {
	o := orm.NewOrm()

	user1 := User{Email: beego.AppConfig.String("user1")}


	err1 := o.Read(&user1, "Email")
	if err1 == nil  {
		quiz1, _ := activateQuizForUser(user1)
		if quiz1 != nil {
			quizList = append(quizList, quiz1)
		}
	}

	user2 := User{Email: beego.AppConfig.String("user2")}
	err2 := o.Read(&user2, "Email")
	if err2 == nil  {

		quiz2, _ := activateQuizForUser(user2)

		if quiz2 != nil {
			quizList = append(quizList, quiz2)
		}
	}

	if len(quizList) == 0 {
		return nil
	}

	return quizList

}

func GetQuizViewData() (data *QuizViewData, err error){

	o := orm.NewOrm()
	user1 := User{Email: beego.AppConfig.String("user1")}
	user2 := User{Email: beego.AppConfig.String("user2")}

	err = o.Read(&user1, "Email")
	if err != nil {
		beego.Warn(err)
		return nil, err
	}

	err = o.Read(&user2, "Email")
	if err != nil {
		beego.Warn(err)
		return nil, err
	}


	quizList1 , err  := getProgressBarData(user1)
	if err != nil {
		beego.Warn(err)
		return nil, err
	}

	quizList2 , err  := getProgressBarData(user2)
	if err != nil {
		beego.Warn(err)
		return nil, err
	}

	data = new(QuizViewData)
	data.Pb1 = prepareProgressBarData(quizList1)
	data.Pb2 = prepareProgressBarData(quizList2)

	return data,nil
}






