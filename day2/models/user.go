package models

import (
	"agmc/day2/config"
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"agmc/day2/entity"
	"fmt"
)

type UserModel struct {
	db config.DatabaseConfig
}

func (e *UserModel) GetAllUser() (*[]entity.User, error) {
	var users []entity.User

	err := e.db.GetDatabaseConfig().Find(&users).Error

	if err != nil {
		fmt.Printf("[userModel.GetUsers] error execute query %v \n", err)
		return nil, fmt.Errorf("failed get all users")
	}

	return &users, nil
}

func (e *UserModel) GetUserByID(id string) (*entity.User, error) {
	var user entity.User

	err := e.db.GetDatabaseConfig().Where("id = ?", id).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("No content found")
		}
		fmt.Printf("[userModel.GetUserByID] error execute query %v \n", err)
		return nil, fmt.Errorf("failed get user data")
	}

	return &user, nil
}

func (e *UserModel) AddUser(user *entity.User) (*entity.User, error) {

	u2 := uuid.NewV4()

	user.ID = u2

	err := e.db.GetDatabaseConfig().Save(&user).Error

	if err != nil {
		fmt.Printf("[userModel.AddUser] error execute query %v \n", err)
		return nil, fmt.Errorf("failed add new user")
	}

	return user, nil
}

func (e *UserModel) UpdateUser(id string, user *entity.User) (*entity.User, error) {

	userData, err := e.GetUserByID(id)

	if err != nil {
		return nil, fmt.Errorf("Failed to update user with id " + id)
	}

	err = e.db.GetDatabaseConfig().Model(&userData).Where("id = ?", id).Updates(entity.User{Name: user.Name, Email: user.Email, Password: user.Password}).Error

	if err != nil {
		fmt.Printf("[userModel.AddUser] error execute query %v \n", err)
		return nil, fmt.Errorf("Failed to update user with id " + id)
	}

	return userData, nil
}

func (e *UserModel) DeleteUser(id string) error {

	userData, err := e.GetUserByID(id)

	if err != nil {
		return fmt.Errorf("Failed to delete user with id " + id)
	}

	err = e.db.GetDatabaseConfig().Delete(userData).Error

	if err != nil {
		fmt.Printf("[userModel.AddUser] error execute query %v \n", err)
		return fmt.Errorf("Failed to delete user with id " + id)
	}

	return nil
}
