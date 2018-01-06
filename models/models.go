package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"beego-ripple/utils"
)

const QUIZ_STATUS_NEW  = 0
const QUIZ_STATUS_OK  = 1
const QUIZ_STATUS_FAIL  = 2

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
	fmt.Println("Init Model")
	orm.RegisterModel(new(User), new(Quiz))
}

func InitDb() {

	o := orm.NewOrm()

	cnt, err := o.QueryTable("user").Count() // SELECT COUNT(*) FROM USER
	fmt.Printf("Count Num: %s, %s", cnt, err)

	//if (cnt > 0) {
	//	fmt.Println("Records found");
	//	return
	//}

	// Error.
	err = orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
		return
	}


	user1 := new(User)
	user1.Email = "1example@example.com"
	fmt.Println(o.Insert(user1))
	fillQuizForUser(user1);


	user2 := new(User)
	user2.Email = "2example@example.com"
	fmt.Println(o.Insert(user2))
	fillQuizForUser(user2);

	return
}


func FindNewQuiz(id int, slug string) (quiz Quiz, err error) {

	o := orm.NewOrm()

	qs := o.QueryTable("quiz").
		Filter("id", id).
		Filter("slug", slug).
		Filter("status", QUIZ_STATUS_NEW).RelatedSel()


	err = qs.One(&quiz)

	if err != nil {
		fmt.Println(err)
		return quiz, err
	}

	fmt.Println(quiz)
	fmt.Println(quiz.User)

	return quiz, nil
}

func fillQuizForUser(user * User)  {
	o := orm.NewOrm()

	for i := 0; i < 6; i++ {

		quiz := new(Quiz)
		quiz.User = user
		quiz.Status = QUIZ_STATUS_NEW
		quiz.Slug = utils.RandomString(10)
		//quiz.Step = 1 + 2 * i
		quiz.Step = i+1

		fmt.Println(o.Insert(quiz))
	}

}





