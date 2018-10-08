package repositories

import (
	"errors"
	"github.com/Muart-C/webcbo/models"
)

//CreateUser Method definition
func CreateUserRepository(userName, email, firstName, lastName string) (*models.User, error) {
	var user models.User
	user.UserName = userName
	user.Email = email
	user.FirstName = firstName
	user.LastName = lastName
	//import db instance to continue

	result := models.DB.Create(&user)

	if result != nil {
		return &models.User{}, nil
	}
	return nil, errors.New("Trouble adding a new user")
}

//FetchUsers Method definition
func FetchUsersRepository() (*[]models.User, error) {
	var users []models.User
	models.DB.Find(&users)
	if len(users) > 0 {
		return &users, nil
	}
	return nil, errors.New("You have no user records")
}

//FetchUser Method definition
func FetchUserRepository(id int) (*models.User, error) {
	var user models.User
	models.DB.Where("id= ?", id).Find(&user)
	if user.ID == id {
		return &models.User{}, nil
	}
	return nil, errors.New("No such user exist")
}

//UpdateUser Method definition
func UpdateUserRepository(userName, email, firstName, lastName string, id int) (*models.User, error) {
	var user models.User
	//user, err := FetchUserRepository(id)
	models.DB.Where("id= ?", id).Find(&user)
	if user.ID == id {
		user.UserName = userName
		user.Email = email
		user.FirstName = firstName
		user.LastName = lastName
		models.DB.Save(&user)
	}
	return &models.User{}, nil
}

//DeleteUser Method definition
func DeleteUserRepository(id int) error { //RUN TESTS
	user, err := r.FetchUserRepository(id)
	if err != nil {
		return errors.New("An error occured while fetching the user")
	}
	models.DB.Delete(&user)
	return nil
}

//CreateEmployee method definition
func CreateEmployeeRepository(user models.User, employeeName string) (*models.Employee, error) {
	var employee models.Employee
	employee.EmployeeName = employeeName
	employee.EmployeeUserID = user.ID
	result := models.DB.Create(&employee)

	if result == nil {
		return nil, errors.New("Error adding Employee")
	}
	return &models.Employee{}, nil
}

//FetchEmployee method definition
func FetchEmployeeRepository(id int) (*models.Employee, error) {
	var employee models.Employee
	models.DB.Where("id=?", id).Find(&employee)
	if employee.ID == id {
		return &models.Employee{}, nil
	}
	return nil, errors.New("Employee with that id not found")
}

//FetchEmployees method definition
func FetchEmployeesRepository() (*models.Employee, error) {
	var employees []models.Employee
	models.DB.Find(&employees)
	if len(employees) > 0 {
		return &models.Employee{}, nil
	}
	return nil, errors.New("No employees were found")
}

//UpdateEmployee method definition
func UpdateEmployeeRepository(employee1 string, id int) (*models.Employee, error) {
	var employee models.Employee
	//employee,err := FetchEmployeeRepository(id)
	models.DB.Where("id=?", id).Find(&employee)
	if employee.ID == id {

		employee.EmployeeName = employee1
		models.DB.Save(employee)

	}
	return &models.Employee{}, nil
}

//DeleteEmployee method definition
func DeleteEmployeeRepository(id int) error {
	employee, err := r.FetchEmployeeRepository(id) //RUN TESTS
	if err != nil {
		return errors.New("An error occured while deleting an employee")
	}
	models.DB.Delete(&employee)
	return nil
}

//CreateRole Method definition
func CreateRoleRepository(roleName string) (*models.Role, error) {
	var role models.Role
	role.RoleName = roleName

	result := models.DB.Create(&role)
	if result == nil {
		return nil, errors.New("Error adding a new role")
	}
	return &models.Role{}, nil

}

//FetchRole method definition
func FetchRoleRepository(id int) (*models.Role, error) {
	var role models.Role
	models.DB.Where("id=?", id).Find(&role)

	if role.ID == id {
		return &models.Role{}, nil
	}
	return nil, errors.New("Role not found")
}

//FetchRoles method definition
func FetchRolesRepository() (*models.Role, error) {
	var roles []models.Role
	models.DB.Find(&roles)
	if len(roles) > 0 {
		return &models.Role{}, nil
	}
	return nil, errors.New("No roles found")
}

//UpdateRole method definition
func UpdateRoleRepository(role1 string, id int) (*models.Role, error) {
	role, err := FetchRoleRepository(id)
	if err != nil {
		return nil, errors.New("Role with that id not found")
	}
	role.RoleName = role1
	models.DB.Save(role)

	return &models.Role{}, nil
}

//DeleteRole method definition
func DeleteRoleRepository(id int) error {
	role, err := FetchRoleRepository(id)
	if err != nil {
		return errors.New("Role with that id not found")
	}
	models.DB.Delete(&role)
	return nil
}
