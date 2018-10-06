package models

import "time"

type Task struct {
	BaseModel
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
}

type TaskStatus struct {
	BaseModel
	TaskStatus string `gorm:"type:varchar(40)"`
	TaskPriority string `gorm:"type:varchar(10)"`
}

type PreviousTask struct {
	BaseModel
}

type Activity struct {
	BaseModel
	ActivityName string
	ActivityPlannedBudget float64 `gorm:type:decimal(10,2)"`
	ActivityActualBudget float64 `gorm:type:decimal(10,2)"`
	ActivityPlannedStartDate time.Time `gorm:"type:timestamp" json:"startedOn"`
	ActivityPlannedEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	ActivityActualStartDate time.Time `gorm:"type:timestamp" json:"StartedOn"`
	ActivityActualEndDate time.Time `gorm:"type:timestamp" json:"EndedOn"`
	//Join one to one between activity status and activity
}

type ActivityStatus struct {
	BaseModel
	ActivityStatus string `gorm:"type:varchar(12)"`
}


type PreviousActivity struct {
	BaseModel
}


