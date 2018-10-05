package models

type Project struct {
	BaseModel
	ProjectName string `gorm:"not null"`
	ProjectDescription interface{} `gorm:"type:varchar(255);not null"`
	ProjectBudget float64 `gorm:"null"`
	ProjectLaborCost float64 `gorm`
}
