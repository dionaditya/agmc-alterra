package entity

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primary_key"`
	Name     string    `gorm:"not null" json:"Name" validate:"required"`
	Email    string    `gorm:"not null" json:"Email" validate:"required"`
	Password string    `gorm:"not null" json:"Password,omitempty" validate:"required"`
}
