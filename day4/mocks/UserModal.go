package mocks

import (
	"agmc/day2/entity"
	"fmt"
)

func GetAllUserModelMock(isError bool) (*[]entity.User, error) {
	var users []entity.User

	if isError {
		return nil, fmt.Errorf("failed get all users")
	}

	return &users, nil
}
