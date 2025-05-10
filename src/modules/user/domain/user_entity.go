package domain_user

import (
	domain_belt "bjj-tracker/src/modules/belt/domain"
	domain_training "bjj-tracker/src/modules/training/domain"
	"time"
)

type User struct {
	ID               string                            `gorm:"primarykey"`
	Name             string                            `json:"name" gorm:"not null"`
	Username         string                            `json:"username" gorm:"not null;unique"`
	Avatar           string                            `json:"avatar" gorm:"not null"`
	Email            string                            `json:"email" gorm:"not null;unique"`
	Password         string                            `json:"password" gorm:"not null"`
	BeltProgress     []domain_belt.BeltProgress        `gorm:"foreignKey:UserID"`
	TrainingSessions []domain_training.TrainingSession `gorm:"foreignKey:UserID"`
	CreatedAt        time.Time                         `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time                         `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
