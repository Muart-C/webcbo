package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Project struct {
	gorm.Model
	ProjectStatusID         uint
	ProjectName             string  `gorm:"not null"`
	ProjectDescription      string  `gorm:"type:text; not null"`
	ProjectBudget           float64 `gorm:"null"`
	ProjectLaborCost        float64 `gorm:"type:decimal(10,2)"`
	ProjectMaterialCost     float64 `gorm:type:decimal(10,2)"`
	ProjectPlannedStartDate time.Time
	ProjectPlannedEndDate   time.Time
	ProjectActualStartDate  time.Time
	ProjectActualEndDate    time.Time
	ProjectTotalHours       float64          `gorm:"type:decimal(10,2)"`
	ProjectManager          []ProjectManager `gorm:"foreignkey:ProjectManagerProjectID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type ProjectStatus struct {
	gorm.Model
	ProjectStatus   string `gorm:"type:varchar(12)"`
	ProjectIsActive bool
	Project         Project `gorm:"foreignkey:ProjectStatusID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type Milestone struct {
	gorm.Model
	MilestoneStatusID     uint
	MilestoneName         string  `gorm:"not null"`
	MilestoneDeliverables string  `gorm:"type:text; not null"`
	MilestoneTotalHours   float64 `gorm:"type:decimal(10,2)"`
	MilestoneDueDate      time.Time
	MilestoneAchievedDate time.Time
}

type MilestoneStatus struct {
	gorm.Model
	MilestoneStatus string    `gorm:"type:varchar(12)"`
	Milestone       Milestone `gorm:"foreignkey:MilestoneStatusID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}
