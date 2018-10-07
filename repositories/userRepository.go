package repositories

import (
	"errors"
	"github.com/Muart-C/webcbo/models"
	"github.com/satori/go.uuid"
)

//CreateUser Method definition
func CreateUser(userName, email, firstName, lastName string) (*models.User, error) {
	var user models.User
	user.ID = uuid.Must(uuid.NewV4())
	user.UserName = userName
	user.Email = email
	user.FirstName = firstName
	user.LastName = lastName
	//import db instance to continue

	models.DB.Create(&user)

	//if user.ID != 0 {
	//	return &models.User{}, nil
	//}
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
func FetchUser(id uuid.UUID) (*models.User, error) {
	var user models.User
	models.DB.Where("id= ?", id).Find(&user)
	if user.ID == id {
		return &models.User{}, nil
	}
	return nil, errors.New("No such user exist")
}

//UpdateUser Method definition
func UpdateUser(userName, email, firstName, lastName string, id uuid.UUID) (*models.User, error) {
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
func DeleteUser(id uuid.UUID) error {
	user, err := FetchUser(id)
	if err != nil {
		return errors.New("An error occured while updating the user details")
	}
	models.DB.Delete(&user)
	return nil
}

//AddRole Method definition
func AddRole(roleName string) (*models.Role, error) {
	var role models.Role
	role.ID = uuid.Must(uuid.NewV4())
	role.RoleName = roleName

	models.DB.Create(&role)
	return &models.Role{}, nil

}
