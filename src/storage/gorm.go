package storage

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func init() {

}

type Issue struct {
	Id        int
	Content   string
	Author    string `sql:"not null`
	Discusses []Discuss
	CreatedAt time.Time
}

type Discuss struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	IssueId   int    `sql:"index"`
	CreatedAt time.Time
}

var Dbg *gorm.DB

func init() {
	fmt.Println("init gorm start")
	// var err error
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// Dbg, err = gorm.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }

	// Dbg.AutoMigrate(&Issue{}, &Discuss{})

	// fmt.Println("init gorm done.")
}

func GORM() {
	fmt.Println("GORM start")
	issue := Issue{Content: "Hello World Issue!", Author: "Sau Sheong"}
	fmt.Println(issue)

	Dbg.Create(&issue)
	fmt.Println(issue)

	discuss := Discuss{Content: "Good discuss", Author: "Joe"}
	Dbg.Model(&issue).Association("Discusses").Append(discuss)
	fmt.Println(issue)

	var readIssue Issue
	Dbg.Where("author=$1", "Sau Sheong").First(&readIssue)
	var discusses []Discuss
	Dbg.Model(&readIssue).Related(&discusses)
	fmt.Println(discusses)

}
