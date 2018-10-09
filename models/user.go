package models

import (
	"errors"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/genproto/googleapis/type/date"
)

//User model
type User struct {
	BaseModel
	UserName       string           `gorm:"not null"`
	Email          string           `gorm:"unique; not null"`
	FirstName      string           `gorm:"type:varchar(30)"`
	LastName       string           `gorm:"type:varchar(30)"`
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
	EmployeeUserID int
	EmployeeName   string       `gorm:"varchar(30)"`
	Assigned       []Assigned   `gorm:"foreignkey:AssignedRoleID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	TeamMember     []TeamMember `gorm:"foreignkey:TeamMemberEmployeeID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
	Hours          Hours        `gorm:"foreignkey:HoursEmployeeID;association_autoupdate:false;association_autocreate:false;association_save_reference:false"`
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
	AssignedOn      date.Date
	AssignedAt      timestamp.Timestamp
	HoursCompleted  float64 `gorm:"type:decimal(10,2)"`
	WorkCompleted   string  `gorm:"type:text"`
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
	//Join project_id,project_status_id,user_id,client_id
	ProjectManagerUserID    int
	ProjectManagerProjectID int
	ProjectManagerClientID  int
}

//Assigned model
type Assigned struct {
	BaseModel
	//Join activity,employee and role(All have 1 to many relationship)
	AssignedActivityID int
	AssignedEmployeeID int
	AssignedRoleID     int
}

//CreateUser Method definition
func CreateUser(userName, email, firstName, lastName string) (*User, error) {
	var user User
	user.UserName = userName
	user.Email = email
	user.FirstName = firstName
	user.LastName = lastName
	result := DB.Create(&user)

	if result != nil {
		return &user, nil
	}
	return nil, errors.New("Trouble adding a new user")
}

//FetchUsers Method definition
func FetchUsers() (*[]User, error) {
	var users []User
	DB.Find(&users)
	if len(users) > 0 {
		return &users, nil
	}
	return nil, errors.New("You have no user records")
}

//FetchUser Method definition
func FetchUser(id int) (*User, error) {
	var user User
	DB.Where("id= ?", id).Find(&user)
	if user.ID == id {
		return &user, nil
	}
	return nil, errors.New("No such user exist")
}

//UpdateUser Method definition
func UpdateUser(userName, email, firstName, lastName string, id int) (*User, error) {
	var user User
	//user, err := FetchUser(id)
	DB.Where("id= ?", id).Find(&user)
	if user.ID == id {
		user.UserName = userName
		user.Email = email
		user.FirstName = firstName
		user.LastName = lastName
		DB.Save(&user)
	}
	return &user, nil
}

//DeleteUser Method definition
func DeleteUser(id int) error { //RUN TESTS
	user, err := FetchUser(id)
	if err != nil {
		return errors.New("An error occured while fetching the user")
	}
	DB.Delete(&user)
	return nil
}

//CreateEmployee method definition
func CreateEmployee(user User, employeeName string) (*Employee, error) {
	var employee Employee
	employee.EmployeeName = employeeName
	employee.EmployeeUserID = user.ID
	result := DB.Create(&employee)

	if result == nil {
		return nil, errors.New("Error adding Employee")
	}
	return &employee, nil
}

//FetchEmployee method definition
func FetchEmployee(id int) (*Employee, error) {
	var employee Employee
	DB.Where("id=?", id).Find(&employee)
	if employee.ID == id {
		return &employee, nil
	}
	return nil, errors.New("Employee with that id not found")
}

//FetchEmployees method definition
func FetchEmployees() (*[]Employee, error) {
	var employees []Employee
	DB.Find(&employees)
	if len(employees) > 0 {
		return &employees, nil
	}
	return nil, errors.New("No employees were found")
}

//UpdateEmployee method definition
func UpdateEmployee(employee1 string, id int) (*Employee, error) {
	var employee Employee
	//employee,err := FetchEmployee(id)
	DB.Where("id=?", id).Find(&employee)
	if employee.ID == id {

		employee.EmployeeName = employee1
		DB.Save(employee)

	}
	return &employee, nil
}

//DeleteEmployee method definition
func DeleteEmployee(id int) error {
	employee, err := FetchEmployee(id) //RUN TESTS
	if err != nil {
		return errors.New("An error occured while deleting an employee")
	}
	DB.Delete(&employee)
	return nil
}

//CreateRole Method definition
func CreateRole(roleName string) (*Role, error) {
	var role Role
	role.RoleName = roleName

	result := DB.Create(&role)
	if result == nil {
		return nil, errors.New("Error adding a new role")
	}
	return &role, nil

}

//FetchRole method definition
func FetchRole(id int) (*Role, error) {
	var role Role
	DB.Where("id=?", id).Find(&role)

	if role.ID == id {
		return &role, nil
	}
	return nil, errors.New("Role not found")
}

//FetchRoles method definition
func FetchRoles() (*[]Role, error) {
	var roles []Role
	DB.Find(&roles)
	if len(roles) > 0 {
		return &roles, nil
	}
	return nil, errors.New("No roles found")
}

//UpdateRole method definition
func UpdateRole(role1 string, id int) (*Role, error) {
	role, err := FetchRole(id)
	if err != nil {
		return nil, errors.New("Role with that id not found")
	}
	role.RoleName = role1
	DB.Save(*role)

	return *&role, nil
}

//DeleteRole method definition
func DeleteRole(id int) error {
	role, err := FetchRole(id)
	if err != nil {
		return errors.New("Role with that id not found")
	}
	DB.Delete(&role)
	return nil
}
