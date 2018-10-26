package models

//Task model
type Task struct {
	BaseModel
	TaskProjectID        int
	TaskName             string
	TaskInstructions     string `gorm:"type:text"`
	TaskTotalHours       float64
	TaskPlannedBudget    float64
	TaskActualBudget     float64
	TaskPlannedStartDate string
	TaskPlannedEndDate   string
	TaskActualStartDate  string
	TaskActualEndDate    string
	TaskStatus           string         `gorm:"type:varchar(40)"`
	TaskPriority         string         `gorm:"type:varchar(10)"` //figure out how to work with enum for a drop down list
	Activity             []Activity     `gorm:"foreignkey:TaskActivityID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	PreviousTask         []PreviousTask `gorm:"foreignkey:PreviousTaskID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//PreviousTask  model
type PreviousTask struct {
	BaseModel
	PreviousTaskID int
}

//Activity model
type Activity struct {
	BaseModel
	TaskActivityID           int
	ActivityName             string
	ActivityPlannedBudget    float64
	ActivityActualBudget     float64
	ActivityPlannedStartDate string
	ActivityPlannedEndDate   string
	ActivityActualStartDate  string
	ActivityActualEndDate    string
	ActivityStatus           string             `gorm:"type:varchar(12)"`
	Task                     []Task             `gorm:"foreignkey:TaskActivityID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	PreviousActivity         []PreviousActivity `gorm:"foreignkey:PreviousActivityID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	Assigned                 []Assigned         `gorm:"foreignkey:AssignedActivityID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//PreviousActivity model
type PreviousActivity struct {
	BaseModel
	PreviousActivityID int
}
