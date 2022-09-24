package usecase

import (
	"agmc/day2/internal/domains"
	"agmc/day2/internal/middleware"
	"agmc/day2/pkg/utils"
)

type UserUseCase struct {
	userRepo       domains.UserRepository
	authMiddleware middleware.AuthMiddleware
}

func NewUserUseCase(u domains.UserRepository) domains.UserUseCase {
	return &UserUseCase{
		userRepo: u,
	}
}

func (userUseCase *UserUseCase) GetAllUser() (*[]domains.User, error) {
	users, err := userUseCase.userRepo.GetAllUserRepository()

	return users, err
}

func (userUseCase *UserUseCase) GetUserByID(id string) (*domains.User, error) {
	users, err := userUseCase.userRepo.GetUserByIDRepository(id)

	return users, err
}

func (userUseCase *UserUseCase) AddUser(userPayload *domains.User) (*domains.User, error) {

	userData, err := userUseCase.userRepo.AddUserRepository(userPayload)

	return userData, err
}

func (userUseCase *UserUseCase) UpdateUser(id string, payload *domains.User) (*domains.User, error) {

	user, err := userUseCase.userRepo.UpdateUserRepository(id, payload)

	return user, err
}

func (userUseCase *UserUseCase) DeleteUser(id string) error {
	err := userUseCase.userRepo.DeleteUserRepository(id)

	return err
}

func (userUseCase *UserUseCase) SignIn(email string, password string, privateKey []byte) (string, error) {
	var user *domains.User

	err := userUseCase.userRepo.GetUserByEmailRepository(email).First(&user).Error

	if err != nil {
		return "", err
	}

	err = utils.CompareHashAndPassword(user.Password, password)

	if err != nil {
		return "", err
	}

	token, err := userUseCase.authMiddleware.GenerateToken(privateKey, user.Email)

	if err != nil {
		return "", err
	}

	return token, err
}
