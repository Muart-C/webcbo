package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
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

func init()  {
	env := godotenv.Load() //loads env variables from .env
	if env != nil {
		log.Fatal(env)
	}
	db = Connect()
	db.AutoMigrate(User{},Project{},Task{},Activity{},Assigned{},Client{},Employee{},Hours{},Milestone{},MilestoneStatus{},PreviousActivity{},PreviousTask{},ProjectManager{},ProjectStatus{},Role{},Task{},TaskStatus{},Team{},TeamMember{})

}



//Database connection
func Connect() (db *gorm.DB) {
	DBHOST := os.Getenv("DBHOST")
	DBNAME := os.Getenv("DBNAME")
	DBPASS := os.Getenv("DBPASS")
	DBSSL:= os.Getenv("DBSSLMODE")
	db, connError := gorm.Open("postgres","host="+DBHOST+"port=5432 dbname="+DBNAME+"dbpass="+DBPASS+"ssl="+DBSSL)
	if connError != nil {
		fmt.Println("DB Connection ERROR")
		log.Fatal(connError)
	}
	return db
}
