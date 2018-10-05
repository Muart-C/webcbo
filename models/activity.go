package models

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
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
	Automigrate(values interface{}) *gorm.DB
}

func init()  {
	env := godotenv.Load() //loads env variables from .env
	if env != nil {
		log.Fatal(env)
	}
	db = Connect()
	db.Automigrate(User{},Project{},Task{},Activity{},Assigned{},)

}
