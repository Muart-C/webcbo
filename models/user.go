package models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/genproto/googleapis/type/date"
)

type User struct {
	BaseModel
	Refer int
	UserName string `gorm:"not null"`
	Email string `gorm:"unique; not null"`
	FirstName string `gorm:"type:varchar(30)"`
	LastName string `gorm:"type:varchar(30)"`
}

type Role struct {
	BaseModel
	Refer int
	RoleName string `gorm:"not null"`
}

type Team struct {
	BaseModel
	Refer int
	TeamName string `gorm:"not null"`
}

type Employee struct {
	BaseModel
	Refer int
	EmployeeName string `gorm:"varchar(30)"`
}

type TeamMember struct {
	BaseModel
	//Join not implemented role,employee,team
	Employee Employee `gorm:"association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	Team Team `gorm:"association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	Role Role `gorm:"association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type Hours struct {
	BaseModel
	AssignedOn date.Date
	AssignedAt timestamp.Timestamp
	HoursCompleted float64 `gorm:"type:decimal(10,2)"`
	WorkCompleted string `gorm:"type:text"`
	//Implement foreign key from employee table one to one relationship
}

type Client struct {
	BaseModel
	ClientName string `gorm:"not null"`
	ClientDescription string `gorm:"type:text"`
	ClientCounty string `gorm:"type:varchar(30)"`
	ClientIndustrySector string `gorm:"type:varchar(128)"`
	ClientCity string `gorm:"type:varchar(128)"`
	ClientPhone string `gorm:"type:varchar(128)"`
	ClientZip string `gorm:"type:varchar(128)"`
}

type ProjectManager struct {
	BaseModel
	//Join project_id,project_status_id,user_id,client_id
}

type Assigned struct {
	BaseModel
	//Join activity,employee and role(All have 1 to many relationship)
}


