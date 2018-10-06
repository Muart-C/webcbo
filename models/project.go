package models

import "time"

type Project struct {
	BaseModel
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
	//Implement Join one to one between project status and project
}

type ProjectStatus struct {
	BaseModel
	ProjectStatus string `gorm:"type:varchar(12)"`
	ProjectIsActive bool
}

type Milestone struct {
	BaseModel
	MilestoneName string `gorm:"not null"`
	MilestoneDeliverables string `gorm:"type:text; not null"`
	MilestoneTotalHours float64 `gorm:"type:decimal(10,2)"`
	MilestoneDueDate time.Time `gorm:"type:timestamp" json:"DueOn"`
	MilestoneAchievedDate time.Time `gorm:"type:timestamp" json:"AchievedOn"`
	//Implement Join one to one between milestone status and milestone
}

type MilestoneStatus struct {
	BaseModel
	MilestoneStatus string `gorm:"type:varchar(12)"`
}
