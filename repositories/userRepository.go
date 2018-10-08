package repositories

import (
	"errors"
	"github.com/Muart-C/webcbo/models"
)

//CreateUser Method definition
func CreateUser(userName, email, firstName, lastName string) (*models.User, error) {
	var user models.User
	user.UserName = userName
	user.Email = email
	user.FirstName = firstName
	user.LastName = lastName
	//import db instance to continue

	result := models.DB.Create(&user)

	if result != nil{
		return &models.User{}, nil
	}
	return nil, errors.New("Trouble adding a new user")
}

//FetchUsers Method definition
func FetchUsers() (*[]models.User, error) {
	var users []models.User
	models.DB.Find(&users)
	if len(users) > 0 {
		return &users, nil
	}
	return nil, errors.New("You have no user records")
}

//FetchUser Method definition
func FetchUser(id int) (*models.User, error) {
	var user models.User
	models.DB.Where("id= ?", id).Find(&user)
	if user.ID == id {
		return &models.User{}, nil
	}
	return nil, errors.New("No such user exist")
}

//UpdateUser Method definition
func UpdateUser(userName, email, firstName, lastName string, id int) (*models.User, error) {
	user, err := FetchUser(id)
	if err != nil {
		return nil, errors.New("An error occured while updating the user details")
	}
	user.UserName = userName
	user.Email = email
	user.FirstName = firstName
	user.LastName = lastName
	models.DB.Save(&user)
	return &models.User{}, nil
}

//DeleteUser Method definition
func DeleteUser(id int) error {
	user, err := FetchUser(id)
	if err != nil {
		return errors.New("An error occured while fetching the user")
	}
	models.DB.Delete(&user)
	return nil
}


//CreateEmployee method definition
func CreateEmployee(user models.User, employeeName string)(*models.Employee, error)  {
	var employee models.Employee
	employee.EmployeeName = employeeName
	employee.EmployeeUserID = user.ID
	result := models.DB.Create(&employee)

	if  result == nil{
		return nil, errors.New("Error adding Employee")
	}
	return &models.Employee{}, nil
}

//FetchEmployee method definition
func FetchEmployee(id int) (*models.Employee, error)  {
	var employee models.Employee
	models.DB.Where("id=?",id).Find(&employee)
	if employee.ID == id {
		return &models.Employee{}, nil
	}
	return nil, errors.New("Employee with that id not found")
}

//FetchEmployees method definition
func FetchEmployees()(*models.Employee,error)  {
	var employees []models.Employee
	models.DB.Find(&employees)
	if len(employees)>0 {
		return &models.Employee{}, nil
	}
	return nil, errors.New("No employees were found")
}

//UpdateEmployee method definition
func UpdateEmployee(employee1 string, id int)(*models.Employee, error)  {
	employee,err := FetchEmployee(id)
	if err != nil {
		return nil, errors.New("Error retrieving the employee")
	}
	employee.EmployeeName = employee1
	models.DB.Save(&employee)
	////////
}
//CreateRole Method definition
func CreateRole(roleName string) (*models.Role, error) {
	var role models.Role
	role.RoleName = roleName

	result := models.DB.Create(&role)
	if result == nil {
		return nil, errors.New("Error adding a new role")
	}
	return &models.Role{}, nil

}

//FetchRole method definition
func FetchRole(id int) (*models.Role, error) {
	var role models.Role
	models.DB.Where("id=?",id).Find(&role)

	if role.ID == id {
		return &models.Role{}, nil
	}
	return nil, errors.New("Role not found")
}

//FetchRoles method definition
func FetchRoles()(*models.Role,error)  {
	var roles []models.Role
	models.DB.Find(&roles)
	if len(roles) > 0 {
		return &models.Role{}, nil
	}
	return nil, errors.New("No roles found")
}

//UpdateRole method definition
func UpdateRole(role1 string, id int)(*models.Role, error)  {
	role, err := FetchRole(id)
	if err != nil {
		return nil, errors.New("Role with that id not found")
	}
	role.RoleName = role1
	models.DB.Save(role)

	return &models.Role{}, nil
}

//DeleteRole method definition
func DeleteRole(id int) error {
	role, err := FetchRole(id)
	if err != nil {
		return  errors.New("Role with that id not found")
	}
	models.DB.Delete(&role)
	return nil
}

