package models

//Project Model definition
type Project struct {
	BaseModel
	ProjectName             string `gorm:"not null"`
	ProjectDescription      string `gorm:"type:text; not null"`
	ProjectBudget           float64
	ProjectLaborCost        float64
	ProjectMaterialCost     float64
	ProjectPlannedStartDate string
	ProjectPlannedEndDate   string
	ProjectActualStartDate  string
	ProjectActualEndDate    string
	ProjectTotalHours       float64
	ProjectStatus           string
	ProjectIsActive         bool
	Task                    []Task           `gorm:"foreignkey:TaskProjectID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	Milestone               []Milestone      `gorm:"foreignkey:MilestoneProjectID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	ProjectManager          []ProjectManager `gorm:"foreignkey:ProjectManagerProjectID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//Milestone Model definition
type Milestone struct {
	BaseModel
	MilestoneProjectID    int
	MilestoneName         string `gorm:"not null"`
	MilestoneDeliverables string `gorm:"type:text; not null"`
	MilestoneTotalHours   float64
	MilestoneDueDate      string
	MilestoneAchievedDate string
	MilestoneStatus       string
}
