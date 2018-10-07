package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	TaskStatusID         uint
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

type TaskStatus struct {
	gorm.Model
	TaskStatus   string `gorm:"type:varchar(40)"`
	TaskPriority string `gorm:"type:varchar(10)"`
	Task         Task   `gorm:"foreignkey:TaskStatusID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type PreviousTask struct {
	gorm.Model
	PreviousTask uint
}

type Activity struct {
	gorm.Model
	ActivityStatusID         uint
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

type ActivityStatus struct {
	gorm.Model
	ActivityStatus string   `gorm:"type:varchar(12)"`
	Activity       Activity `gorm:"foreignkey:ActivityStatusID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type PreviousActivity struct {
	gorm.Model
	PreviousActivityID uint
}
