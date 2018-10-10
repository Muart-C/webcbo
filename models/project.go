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
	ProjectBudget           float64 `gorm:"null"`
	ProjectLaborCost        float64
	ProjectMaterialCost     float64
	ProjectPlannedStartDate time.Time
	ProjectPlannedEndDate   time.Time
	ProjectActualStartDate  time.Time
	ProjectActualEndDate    time.Time
	ProjectTotalHours       float64          `gorm:"type:decimal(10,2)"`
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
	MilestoneTotalHours   float64 `gorm:"type:decimal(10,2)"`
	MilestoneDueDate      time.Time
	MilestoneAchievedDate time.Time
}

//MilestoneStatus Model definition
type MilestoneStatus struct {
	BaseModel
	MilestoneStatus string    `gorm:"type:varchar(12)"`
	Milestone       Milestone `gorm:"foreignkey:MilestoneStatusID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}
