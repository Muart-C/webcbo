package models

import "time"

type Task struct {
	BaseModel
	Refer int
	TaskName string
	TaskInstructions string `gorm:"type:text"`
	TaskTotalHours float64 `gorm:"type:decimal(10,2)"`
	TaskPlannedBudget float64 `gorm:type:decimal(10,2)"`
	TaskActualBudget float64 `gorm:type:decimal(10,2)"`
	TaskPlannedStartDate time.Time `gorm:"type:timestamp" json:"startedOn"`
	TaskPlannedEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	TaskActualStartDate time.Time `gorm:"type:timestamp" json:"StartedOn"`
	TaskActualEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	//Join one to one between task status and task
	TaskStatus TaskStatus `gorm:"association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type TaskStatus struct {
	BaseModel
	Refer int
	TaskStatus string `gorm:"type:varchar(40)"`
	TaskPriority string `gorm:"type:varchar(10)"`
}

type PreviousTask struct {
	BaseModel
	Task Task `gorm:"association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type Activity struct {
	BaseModel
	Refer int
	ActivityName string
	ActivityPlannedBudget float64 `gorm:type:decimal(10,2)"`
	ActivityActualBudget float64 `gorm:type:decimal(10,2)"`
	ActivityPlannedStartDate time.Time `gorm:"type:timestamp" json:"startedOn"`
	ActivityPlannedEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	ActivityActualStartDate time.Time `gorm:"type:timestamp" json:"StartedOn"`
	ActivityActualEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	//Join one to one between activity status and activity
	ActivityStatus ActivityStatus `gorm:"association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

type ActivityStatus struct {
	BaseModel
	Refer int
	ActivityStatus string `gorm:"type:varchar(12)"`
	Activity Activity `gorm:"association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}


type PreviousActivity struct {
	BaseModel
	Activity Activity `gorm:"association_foreignkey:Refer;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}


