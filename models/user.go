package models

//JwtToken authentication token
type JwtToken struct {
	Token string `json:"token"`
}


//User model
type User struct {
	BaseModel
	UserName       string           `gorm:"not null"`
	Email          string           `gorm:"unique; not null"`
	FirstName      string           `gorm:"type:varchar(30)"`
	LastName       string           `gorm:"type:varchar(30)"`
	Password       []byte
	ProjectManager []ProjectManager `gorm:"foreignkey:ProjectManagerUserID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	Employee       []Employee       `gorm:"foreignkey:EmployeeUserID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//Role model
type Role struct {
	BaseModel
	RoleName   string       `gorm:"not null"`
	Assigned   []Assigned   `gorm:"foreignkey:AssignedRoleID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	TeamMember []TeamMember `gorm:"foreignkey:TeamMemberRoleID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//Team model
type Team struct {
	BaseModel
	TeamName   string       `gorm:"not null"`
	TeamMember []TeamMember `gorm:"foreignkey:TeamMemberTeamID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//Employee model
type Employee struct {
	BaseModel
	EmployeeUserID     int
	EmployeeProfession string       `gorm:"varchar(30)"`
	Assigned           []Assigned   `gorm:"foreignkey:AssignedRoleID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	TeamMember         []TeamMember `gorm:"foreignkey:TeamMemberEmployeeID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	Hours              Hours        `gorm:"foreignkey:HoursEmployeeID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//TeamMember model
type TeamMember struct {
	BaseModel
	TeamMemberTeamID     int
	TeamMemberEmployeeID int
	TeamMemberRoleID     int
}

//Hours model
type Hours struct {
	BaseModel
	HoursEmployeeID int
	AssignedOn      string
	AssignedAt      string
	HoursCompleted  float64
	WorkCompleted   string `gorm:"type:text"`
}

//Client model
type Client struct {
	BaseModel
	ClientName           string           `gorm:"not null"`
	ClientDescription    string           `gorm:"type:text"`
	ClientCounty         string           `gorm:"type:varchar(30)"`
	ClientIndustrySector string           `gorm:"type:varchar(128)"`
	ClientCity           string           `gorm:"type:varchar(128)"`
	ClientPhone          string           `gorm:"type:varchar(128)"`
	ClientZip            string           `gorm:"type:varchar(128)"`
	ProjectManager       []ProjectManager `gorm:"foreignkey:ProjectManagerClientID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
}

//ProjectManager model
type ProjectManager struct {
	BaseModel

	ProjectManagerUserID    int
	ProjectManagerProjectID int
	ProjectManagerClientID  int
}

//Assigned model
type Assigned struct {
	BaseModel
	AssignedActivityID int
	AssignedEmployeeID int
	AssignedRoleID     int
}
