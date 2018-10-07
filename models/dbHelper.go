package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db database
)

//interface of functions to be used

type database interface {
	Create(interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	AutoMigrate(values ...interface{}) *gorm.DB
}

func init() {
	db = Connect()
	db.AutoMigrate(User{}, Project{}, Task{}, Activity{}, Assigned{}, Client{}, Employee{}, Hours{}, Milestone{}, MilestoneStatus{}, PreviousActivity{}, PreviousTask{}, ProjectManager{}, ProjectStatus{}, Role{}, Task{}, TaskStatus{}, Team{}, TeamMember{}, ActivityStatus{})

}

//Database connection
func Connect() (db *gorm.DB) {
	db, connError := gorm.Open("mysql", "root:geeky254@tcp(127.0.0.1:3306)/acbo?charset=utf8&parseTime=True")
	if connError != nil {
		fmt.Println("DB Connection ERROR")
		log.Fatal(connError)
	}
	return db
}
