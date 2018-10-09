package repository

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
	result := models.DB.Create(&user)

	if result != nil {
		return &user, nil
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
		return &user, nil
	}
	return nil, errors.New("No such user exist")
}

//UpdateUser Method definition
func UpdateUser(userName, email, firstName, lastName string, id int) (*models.User, error) {
	var user models.User
	//user, err := FetchUser(id)
	models.DB.Where("id= ?", id).Find(&user)
	if user.ID == id {
		user.UserName = userName
		user.Email = email
		user.FirstName = firstName
		user.LastName = lastName
		models.DB.Save(&user)
	}
	return &user, nil
}

//DeleteUser Method definition
func DeleteUser(id int) error { //RUN TESTS
	user, err := FetchUser(id)
	if err != nil {
		return errors.New("An error occured while fetching the user")
	}
	models.DB.Delete(&user)
	return nil
}

//CreateEmployee method definition
func CreateEmployee(user models.User, employeeProfession string) (*models.Employee, error) {
	var employee models.Employee
	employee.EmployeeProfession = employeeProfession
	employee.EmployeeUserID = user.ID
	result := models.DB.Create(&employee)

	if result == nil {
		return nil, errors.New("Error adding Employee")
	}
	return &employee, nil
}

//FetchEmployee method definition
func FetchEmployee(id int) (*models.Employee, error) {
	var employee models.Employee
	models.DB.Where("id=?", id).Find(&employee)
	if employee.ID == id {
		return &employee, nil
	}
	return nil, errors.New("Employee with that id not found")
}

//FetchEmployees method definition
func FetchEmployees() (*[]models.Employee, error) {
	var employees []models.Employee
	models.DB.Find(&employees)
	if len(employees) > 0 {
		return &employees, nil
	}
	return nil, errors.New("No employees were found")
}

//UpdateEmployee method definition
func UpdateEmployee(employeeProfession string, id int) (*models.Employee, error) {
	var employee models.Employee
	//employee,err := FetchEmployee(id)
	models.DB.Where("id=?", id).Find(&employee)
	if employee.ID == id {

		employee.EmployeeProfession = employeeProfession
		models.DB.Save(employee)

	}
	return &employee, nil
}

//DeleteEmployee method definition
func DeleteEmployee(id int) error {
	employee, err := FetchEmployee(id) //RUN TESTS
	if err != nil {
		return errors.New("An error occured while deleting an employee")
	}
	models.DB.Delete(&employee)
	return nil
}

//CreateRole Method definition
func CreateRole(roleName string) (*models.Role, error) {
	var role models.Role
	role.RoleName = roleName

	result := models.DB.Create(&role)
	if result == nil {
		return nil, errors.New("Error adding a new role")
	}
	return &role, nil

}

//FetchRole method definition
func FetchRole(id int) (*models.Role, error) {
	var role models.Role
	models.DB.Where("id=?", id).Find(&role)

	if role.ID == id {
		return &role, nil
	}
	return nil, errors.New("Role not found")
}

//FetchRoles method definition
func FetchRoles() (*[]models.Role, error) {
	var roles []models.Role
	models.DB.Find(&roles)
	if len(roles) > 0 {
		return &roles, nil
	}
	return nil, errors.New("No roles found")
}

//UpdateRole method definition
func UpdateRole(role1 string, id int) (*models.Role, error) {
	role, err := FetchRole(id)
	if err != nil {
		return nil, errors.New("Role with that id not found")
	}
	role.RoleName = role1
	models.DB.Save(*role)

	return *&role, nil
}

//DeleteRole method definition
func DeleteRole(id int) error {
	role, err := FetchRole(id)
	if err != nil {
		return errors.New("Role with that id not found")
	}
	models.DB.Delete(&role)
	return nil
}
