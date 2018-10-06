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


}
