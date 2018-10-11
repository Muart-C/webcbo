package models

import (
	"time"
)

//Project Model definition
type Project struct {

	BaseModel
	ProjectStatusID         int
	ProjectName             string  `gorm:"not null"`
	ProjectDescription      string  `gorm:"type:text; not null"`
	ProjectBudget           float64
	ProjectLaborCost        float64
	ProjectMaterialCost     float64
	ProjectPlannedStartDate string
	ProjectPlannedEndDate   string
	ProjectActualStartDate  string
	ProjectActualEndDate    string
	ProjectTotalHours       float64
	Task []Task `gorm:"foreignkey:TaskProjectID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	ProjectManager          []ProjectManager `gorm:"foreignkey:ProjectManagerProjectID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//ProjectStatus Model definition
type ProjectStatus struct {
	BaseModel
	ProjectStatus   string `gorm:"type:varchar(12)"`
	ProjectIsActive bool
	Project         Project `gorm:"foreignkey:ProjectStatusID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//Milestone Model definition
type Milestone struct {
	BaseModel
	MilestoneStatusID     int
	MilestoneName         string  `gorm:"not null"`
	MilestoneDeliverables string  `gorm:"type:text; not null"`
	MilestoneTotalHours   float64
	MilestoneDueDate      time.Time
	MilestoneAchievedDate time.Time
}

//MilestoneStatus Model definition
type MilestoneStatus struct {
	BaseModel
	MilestoneStatus string    `gorm:"type:varchar(12)"`
	Milestone       Milestone `gorm:"foreignkey:MilestoneStatusID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}
