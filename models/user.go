package models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
	"google.golang.org/genproto/googleapis/type/date"
)

type User struct {
	gorm.Model
	Refer uint
	UserName string `gorm:"not null"`
	Email string `gorm:"unique; not null"`
	FirstName string `gorm:"type:varchar(30)"`
	LastName string `gorm:"type:varchar(30)"`
	ProjectManager []ProjectManager `gorm:"foreignkey:ProjectManagerUserID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type Role struct {
	gorm.Model
	Refer uint
	RoleName string `gorm:"not null"`
	Assigned []Assigned `gorm:"foreignkey:AssignedRoleID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	TeamMember []TeamMember `gorm:"foreignkey:TeamMemberRoleID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type Team struct {
	gorm.Model
	Refer uint
	TeamName string `gorm:"not null"`
	TeamMember []TeamMember `gorm:"foreignkey:TeamMemberTeamID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type Employee struct {
	gorm.Model
	Refer uint
	EmployeeName string `gorm:"varchar(30)"`
	Assigned []Assigned `gorm:"foreignkey:AssignedRoleID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	TeamMember []TeamMember `gorm:"foreignkey:TeamMemberEmployeeID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	Hours Hours `gorm:"foreignkey:EmployeeID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type TeamMember struct {
	gorm.Model
	TeamMemberTeamID uint
	TeamMemberEmployeeID uint
	TeamMemberRoleID uint
}

type Hours struct {
	gorm.Model
	EmployeeID uint
	AssignedOn date.Date
	AssignedAt timestamp.Timestamp
	HoursCompleted float64 `gorm:"type:decimal(10,2)"`
	WorkCompleted string `gorm:"type:text"`
}

type Client struct {
	gorm.Model
	Refer uint
	ClientName string `gorm:"not null"`
	ClientDescription string `gorm:"type:text"`
	ClientCounty string `gorm:"type:varchar(30)"`
	ClientIndustrySector string `gorm:"type:varchar(128)"`
	ClientCity string `gorm:"type:varchar(128)"`
	ClientPhone string `gorm:"type:varchar(128)"`
	ClientZip string `gorm:"type:varchar(128)"`
	ProjectManager  []ProjectManager `gorm:"foreignkey:ProjectManagerClientID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type ProjectManager struct {
	gorm.Model
	//Join project_id,project_status_id,user_id,client_id
	ProjectManagerUserID uint
	ProjectManagerProjectID uint
	ProjectManagerClientID uint
}

type Assigned struct {
	gorm.Model
	//Join activity,employee and role(All have 1 to many relationship)
	AssignedActivityID uint
	AssignedEmployeeID uint
	AssignedRoleID uint
}


