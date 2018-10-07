package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Project struct {
	gorm.Model
	Refer uint
	ProjectStatusID uint
	ProjectName string `gorm:"not null"`
	ProjectDescription string `gorm:"type:text; not null"`
	ProjectBudget float64 `gorm:"null"`
	ProjectLaborCost float64 `gorm:"type:decimal(10,2)"`
	ProjectMaterialCost float64 `gorm:type:decimal(10,2)"`
	ProjectPlannedStartDate time.Time `gorm:"type:timestamp" json:"startedOn"`
	ProjectPlannedEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	ProjectActualStartDate time.Time `gorm:"type:timestamp" json:"StartedOn"`
	ProjectActualEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	ProjectTotalHours float64 `gorm:"type:decimal(10,2)"`
	ProjectManager []ProjectManager `gorm:"foreignkey:ProjectManagerProjectID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type ProjectStatus struct {
	gorm.Model
	Refer int
	ProjectStatus string `gorm:"type:varchar(12)"`
	ProjectIsActive bool
	Project Project `gorm:"foreignkey:ProjectStatusID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type Milestone struct {
	gorm.Model
	MilestoneStatusID uint
	MilestoneName string `gorm:"not null"`
	MilestoneDeliverables string `gorm:"type:text; not null"`
	MilestoneTotalHours float64 `gorm:"type:decimal(10,2)"`
	MilestoneDueDate time.Time `gorm:"type:timestamp" json:"DueOn"`
	MilestoneAchievedDate time.Time `gorm:"type:timestamp" json:"AchievedOn"`
}

type MilestoneStatus struct {
	gorm.Model
	Refer int
	MilestoneStatus string `gorm:"type:varchar(12)"`
	Milestone Milestone `gorm:"foreignkey:MilestoneStatusID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}
