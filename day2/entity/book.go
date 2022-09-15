package entity

type Book struct {
	Name   string `gorm:"not null" json:"Name" validate:"required"`
	Author string `gorm:"not null" json:"Author" validate:"required"`
	Year   int    `gorm:"not null" json:"Year,omitempty" validate:"required"`
}
