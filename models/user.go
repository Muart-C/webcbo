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
	Password       string
	//ProjectManager []ProjectManager
	//Employee       []Employee
}

//Role model
type Role struct {
	BaseModel
	RoleName   string       `gorm:"not null"`
	Assigned   []Assigned
	TeamMember []TeamMember
}

//Team model
type Team struct {
	BaseModel
	TeamName   string       `gorm:"not null"`
	TeamMember []TeamMember
}

//Employee model
type Employee struct {
	BaseModel
	EmployeeUserID int
	EmployeeProfession string       `gorm:"varchar(30)"`
	Assigned           []Assigned
	TeamMember         []TeamMember
	Hours              []Hours
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
	ProjectManager       []ProjectManager
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