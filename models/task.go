package models

import (
	"github.com/satori/go.uuid"
	"time"
)

//Task model
type Task struct {
	BaseModel
	TaskStatusID         uuid.UUID
	TaskName             string
	TaskInstructions     string  `gorm:"type:text"`
	TaskTotalHours       float64 `gorm:"type:decimal(10,2)"`
	TaskPlannedBudget    float64 `gorm:type:decimal(10,2)"`
	TaskActualBudget     float64 `gorm:type:decimal(10,2)"`
	TaskPlannedStartDate time.Time
	TaskPlannedEndDate   time.Time
	TaskActualStartDate  time.Time
	TaskActualEndDate    time.Time
	PreviousTask         []PreviousTask `gorm:"foreignkey:PreviousTaskID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//TaskStatus model
type TaskStatus struct {
	BaseModel
	TaskStatus   string `gorm:"type:varchar(40)"`
	TaskPriority string `gorm:"type:varchar(10)"`
	Task         Task   `gorm:"foreignkey:TaskStatusID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//PreviousTask  model
type PreviousTask struct {
	BaseModel
	PreviousTaskID uuid.UUID
}

//Activity model
type Activity struct {
	BaseModel
	ActivityStatusID         uuid.UUID
	ActivityName             string
	ActivityPlannedBudget    float64 `gorm:type:decimal(10,2)"`
	ActivityActualBudget     float64 `gorm:type:decimal(10,2)"`
	ActivityPlannedStartDate time.Time
	ActivityPlannedEndDate   time.Time
	ActivityActualStartDate  time.Time
	ActivityActualEndDate    time.Time
	PreviousActivity         []PreviousActivity `gorm:"foreignkey:PreviousActivityID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	Assigned                 []Assigned         `gorm:"foreignkey:AssignedActivityID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//ActivityStatus model
type ActivityStatus struct {
	BaseModel
	ActivityStatus string   `gorm:"type:varchar(12)"`
	Activity       Activity `gorm:"foreignkey:ActivityStatusID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//PreviousActivity model
type PreviousActivity struct {
	BaseModel
	PreviousActivityID uuid.UUID
}
