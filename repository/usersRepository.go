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
		models.DB.Save(&employee)

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

//CreateEmployeeHours method definition
func CreateEmployeeHours(employee models.Employee, assignedOn string, assignedAt string, hoursCompleted float64, workCompleted string) (*models.Hours, error) {
	var employeeHours models.Hours
	employeeHours.AssignedOn = assignedOn
	employeeHours.AssignedAt = assignedAt
	employeeHours.HoursEmployeeID = employee.ID
	employeeHours.HoursCompleted = hoursCompleted
	employeeHours.WorkCompleted = workCompleted
	result := models.DB.Create(&employeeHours)

	if result == nil {
		return nil, errors.New("Error adding employee Hours")
	}
	return &employeeHours, nil
}

//FetchEmployeeHours method definition
func FetchEmployeeHours(id int) (*models.Hours, error) {
	var employeeHours models.Hours
	models.DB.Where("id=?", id).Find(&employeeHours)
	if employeeHours.ID == id {
		return &employeeHours, nil
	}
	return nil, errors.New("employeeHours with that id not found")
}

//FetchEmployeeHours method definition
func FetchEmployeesHours() (*[]models.Hours, error) {
	var employeeHours []models.Hours
	models.DB.Find(&employeeHours)
	if len(employeeHours) > 0 {
		return &employeeHours, nil
	}
	return nil, errors.New("No employeeHourss were found")
}

//UpdateEmployeeHours method definition
func UpdateEmployeeHours(assignedOn, assignedAt string, hoursCompleted float64, workCompleted string, id int) (*models.Hours, error) {
	var employeeHours models.Hours
	//employeeHours,err := FetchemployeeHours(id)
	models.DB.Where("id=?", id).Find(&employeeHours)
	if employeeHours.ID == id {
		employeeHours.AssignedAt = assignedAt
		employeeHours.AssignedOn = assignedOn
		employeeHours.HoursCompleted = hoursCompleted
		employeeHours.WorkCompleted = workCompleted
		models.DB.Save(employeeHours)

	}
	return &employeeHours, nil
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

//CreateTeam Method definition
func CreateTeam(teamName string) (*models.Team, error) {
	var team models.Team
	team.TeamName = teamName

	result := models.DB.Create(&team)
	if result == nil {
		return nil, errors.New("Error adding a new Team")
	}
	return &team, nil

}

//FetchTeam method definition
func FetchTeam(id int) (*models.Team, error) {
	var team models.Team
	models.DB.Where("id=?", id).Find(&team)

	if team.ID == id {
		return &team, nil
	}
	return nil, errors.New("Team not found")
}

//FetchTeams method definition
func FetchTeams() (*[]models.Team, error) {
	var teams []models.Team
	models.DB.Find(&teams)
	if len(teams) > 0 {
		return &teams, nil
	}
	return nil, errors.New("No Teams found")
}

//UpdateTeam method definition
func UpdateTeam(team1 string, id int) (*models.Team, error) {
	team, err := FetchTeam(id)
	if err != nil {
		return nil, errors.New("Team with that id not found")
	}
	team.TeamName = team1
	models.DB.Save(*team)

	return *&team, nil
}

//DeleteTeam method definition
func DeleteTeam(id int) error {
	team, err := FetchTeam(id)
	if err != nil {
		return errors.New("Team with that id not found")
	}
	models.DB.Delete(&team)
	return nil
}

//CreateClient Method definition
func CreateClient(clientName, clientDescription, clientCounty, clientIndustrySector, clientCity, clientPhone, clientZip string) (*models.Client, error) {
	var client models.Client
	client.ClientName = clientName
	client.ClientDescription = clientDescription
	client.ClientCounty = clientCounty
	client.ClientIndustrySector = clientIndustrySector
	client.ClientCity = clientCity
	client.ClientPhone = clientPhone
	client.ClientZip = clientZip
	result := models.DB.Create(&client)

	if result != nil {
		return &client, nil
	}
	return nil, errors.New("Trouble adding a new client")
}

//FetchClients Method definition
func FetchClients() (*[]models.Client, error) {
	var clients []models.Client
	models.DB.Find(&clients)
	if len(clients) > 0 {
		return &clients, nil
	}
	return nil, errors.New("You have no client records")
}

//FetchClient Method definition
func FetchClient(id int) (*models.Client, error) {
	var client models.Client
	models.DB.Where("id= ?", id).Find(&client)
	if client.ID == id {
		return &client, nil
	}
	return nil, errors.New("No such client exist")
}

//UpdateClient Method definition
func UpdateClient(clientName, clientDescription, clientCounty, clientIndustrySector, clientCity, clientPhone, clientZip string, id int) (*models.Client, error) {
	var client models.Client
	//client, err := FetchClient(id)
	models.DB.Where("id= ?", id).Find(&client)
	if client.ID == id {
		client.ClientName = clientName
		client.ClientDescription = clientDescription
		client.ClientCounty = clientCounty
		client.ClientIndustrySector = clientIndustrySector
		client.ClientCity = clientCity
		client.ClientPhone = clientPhone
		client.ClientZip = clientZip
		models.DB.Save(&client)
	}
	return &client, nil
}

//DeleteClient Method definition
func DeleteClient(id int) error { //RUN TESTS
	client, err := FetchClient(id)
	if err != nil {
		return errors.New("An error occured while fetching the client")
	}
	models.DB.Delete(&client)
	return nil
}


//CreateProjectManager method definition
func CreateProjectManager(user models.User,project int,client int) (*models.ProjectManager, error) {
	var manager models.ProjectManager
	manager.ProjectManagerUserID= user.ID
	manager.ProjectManagerProjectID = project
	manager.ProjectManagerClientID = client
	result := models.DB.Create(&manager)

	if result == nil {
		return nil, errors.New("Error Assigning project manager")
	}
	return &manager, nil
}

//FetchProjectManager method definition
func FetchProjectManager(id int) (*models.ProjectManager, error) {
	var manager models.ProjectManager
	models.DB.Where("id=?", id).Find(&manager)
	if manager.ID == id {
		return &manager, nil
	}
	return nil, errors.New("Manager with that id not found")
}

//FetchProjectManagers method definition
func FetchProjectManagers() (*[]models.ProjectManager, error) {
	var managers []models.ProjectManager
	models.DB.Find(&managers)
	if len(managers) > 0 {
		return &managers, nil
	}
	return nil, errors.New("No managers were found")
}

//UpdateProjectManager method definition
func UpdateProjectManager(project, client, id int) (*models.ProjectManager, error) {
	var manager models.ProjectManager
	//employee,err := FetchEmployee(id)
	models.DB.Where("id=?", id).Find(&manager)
	if manager.ID == id {
		manager.ProjectManagerProjectID= project
		manager.ProjectManagerProjectID= client
		models.DB.Save(&manager)

	}
	return &manager, nil
}


