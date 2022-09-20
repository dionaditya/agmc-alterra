package repository

import (
	"agmc/day2/pkg/constant"
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	domains "agmc/day2/internal/domains"
	"agmc/day2/pkg/utils"
	"fmt"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepositoryy(Conn *gorm.DB) domains.UserRepository {
	return &userRepository{
		db: Conn,
	}
}

func (e *userRepository) GetAllUserRepository() (*[]domains.User, error) {
	var users []domains.User

	err := e.db.Find(&users).Error

	if err != nil {
		fmt.Printf("[userRepository.GetAllUserRepository] error execute query %v \n", err)
		return nil, fmt.Errorf(constant.E_UNPROCESSABLE_ENTITY)
	}

	return &users, nil
}

func (e *userRepository) GetUserByIDRepository(id string) (*domains.User, error) {
	var user domains.User

	err := e.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf(constant.E_NOT_FOUND)
		}
		fmt.Printf("[userRepository.GetUserByIDRepository] error execute query %v \n", err)
		return nil, fmt.Errorf(constant.E_UNPROCESSABLE_ENTITY)
	}

	return &user, nil
}

func (e *userRepository) AddUserRepository(user *domains.User) (*domains.User, error) {

	u2 := uuid.NewV4()

	user.ID = u2

	result := e.GetUserByEmailRepository(user.Email)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf(constant.E_DUPLICATE)
	}

	hashedPassword, err := utils.GenerateHashPassword(user.Password)

	if err != nil {
		fmt.Printf("[userRepository.AddUserRepository] error execute query %v \n", err)
		return nil, fmt.Errorf(constant.E_UNPROCESSABLE_ENTITY)
	}

	user.Password = string(hashedPassword)

	err = e.db.Save(&user).Error

	if err != nil {
		fmt.Printf("[userRepository.AddUseRepositoryr] error execute query %v \n", err)
		return nil, fmt.Errorf(constant.E_UNPROCESSABLE_ENTITY)
	}

	user.Password = ""

	return user, nil
}

func (e *userRepository) UpdateUserRepository(id string, user *domains.User) (*domains.User, error) {

	userData, err := e.GetUserByIDRepository(id)

	if err != nil {
		return nil, fmt.Errorf(constant.E_UNPROCESSABLE_ENTITY)
	}

	err = e.db.Model(&userData).Where("id = ?", id).Updates(domains.User{Name: user.Name, Email: user.Email, Password: user.Password}).Error

	if err != nil {
		fmt.Printf("[userRepository.updateUserRepository] error execute query %v \n", err)
		return nil, fmt.Errorf(constant.E_UNPROCESSABLE_ENTITY)
	}

	userData.Password = ""

	return userData, nil
}

func (e *userRepository) DeleteUserRepository(id string) error {

	userData, err := e.GetUserByIDRepository(id)

	if err != nil {
		return fmt.Errorf(constant.E_UNPROCESSABLE_ENTITY)
	}

	err = e.db.Delete(userData).Error

	if err != nil {
		fmt.Printf("[userRepository.DeleteUserRepository] error execute query %v \n", err)
		return fmt.Errorf(constant.E_UNPROCESSABLE_ENTITY)
	}

	userData.Password = ""

	return nil
}

func (e *userRepository) GetUserByEmailRepository(email string) *gorm.DB {
	var resultUser domains.User

	result := e.db.Where("email = ?", email).Take(&resultUser)

	return result
}
