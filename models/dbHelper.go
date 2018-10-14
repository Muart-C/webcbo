package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //invoke mysql driver
	"log"
	"time"
)

//DB Object
var (
	DB Database
)

//Database interface method definition
type Database interface {
	Create(interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	AutoMigrate(values ...interface{}) *gorm.DB
}

func init() {

	DB = Connect()
	DB.AutoMigrate(User{}, Project{}, Task{}, Activity{}, Assigned{}, Client{}, Employee{}, Hours{}, Milestone{}, PreviousActivity{}, PreviousTask{}, ProjectManager{}, Role{}, Task{}, Team{}, TeamMember{})

}


//Connect Database connection
func Connect() (db *gorm.DB) {
	db, connError := gorm.Open("mysql", "root:geeky254@tcp(127.0.0.1:3306)/acbo?charset=utf8&parseTime=True")
	if connError != nil {
		fmt.Println("DB Connection ERROR")
		log.Fatal(connError)
	}
	return db
}

//BaseModel model definition
type BaseModel struct {
	ID        int `gorm:"primary_key;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
