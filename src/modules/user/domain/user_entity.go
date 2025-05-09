package domain_user

import (
	domain_belt "bjj-tracker/src/modules/belt_progress/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	BeltProgress []domain_belt.BeltProgress `gorm:"foreignKey:UserID"`
}
