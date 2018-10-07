package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Refer uint
	TaskStatusID uint
	TaskName string
	TaskInstructions string `gorm:"type:text"`
	TaskTotalHours float64 `gorm:"type:decimal(10,2)"`
	TaskPlannedBudget float64 `gorm:type:decimal(10,2)"`
	TaskActualBudget float64 `gorm:type:decimal(10,2)"`
	TaskPlannedStartDate time.Time `gorm:"type:timestamp" json:"startedOn"`
	TaskPlannedEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	TaskActualStartDate time.Time `gorm:"type:timestamp" json:"StartedOn"`
	TaskActualEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	PreviousTask []PreviousTask `gorm:"foreignkey:PreviousTaskID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`

}

type TaskStatus struct {
	gorm.Model
	Refer uint
	TaskStatus string `gorm:"type:varchar(40)"`
	TaskPriority string `gorm:"type:varchar(10)"`
	Task Task `gorm:"foreignkey:TaskStatusID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type PreviousTask struct {
	gorm.Model
	PreviousTask uint
}

type Activity struct {
	gorm.Model
	Refer uint
	ActivityStatusID uint
	ActivityName string
	ActivityPlannedBudget float64 `gorm:type:decimal(10,2)"`
	ActivityActualBudget float64 `gorm:type:decimal(10,2)"`
	ActivityPlannedStartDate time.Time `gorm:"type:timestamp" json:"startedOn"`
	ActivityPlannedEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	ActivityActualStartDate time.Time `gorm:"type:timestamp" json:"StartedOn"`
	ActivityActualEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	PreviousActivity []PreviousActivity `gorm:"foreignkey:PreviousActivityID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	Assigned []Assigned `gorm:"foreignkey:AssignedActivityID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type ActivityStatus struct {
	gorm.Model
	Refer uint
	ActivityStatus string `gorm:"type:varchar(12)"`
	Activity Activity `gorm:"foreignkey:ActivityStatusID;association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}


type PreviousActivity struct {
	gorm.Model
	PreviousActivityID uint
}


