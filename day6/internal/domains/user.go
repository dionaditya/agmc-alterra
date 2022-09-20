package domains

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primary_key"`
	Name     string    `gorm:"not null" json:"Name" validate:"required"`
	Email    string    `gorm:"not null" json:"Email" validate:"required"`
	Password string    `gorm:"size:100" json:"Password"`
}

type UserRepository interface {
	GetAllUserRepository() (*[]User, error)
	GetUserByIDRepository(id string) (*User, error)
	AddUserRepository(user *User) (*User, error)
	UpdateUserRepository(id string, user *User) (*User, error)
	DeleteUserRepository(id string) error
	GetUserByEmailRepository(email string) *gorm.DB
}

type UserUseCase interface {
	GetAllUser() (*[]User, error)
	GetUserByID(id string) (*User, error)
	AddUser(userPayload *User) (*User, error)
	UpdateUser(id string, user *User) (*User, error)
	DeleteUser(id string) error
	SignIn(email string, password string, privateKey []byte) (string, error)
}
