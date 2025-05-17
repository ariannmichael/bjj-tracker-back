package domain_training

import (
	domain_technique "bjj-tracker/src/modules/technique/domain"
	"time"
)

type TrainingSession struct {
	ID         string                       `json:"id" gorm:"primaryKey"`
	UserID     string                       `json:"user_id" gorm:"not null"`
	User       string                       `json:"user" gorm:"not null"`
	Techniques []domain_technique.Technique `json:"technique" gorm:"not null"`
	Duration   int                          `json:"duration" gorm:"not null"` // in minutes
	Notes      string                       `json:"notes" gorm:"not null"`
	CreatedAt  time.Time                    `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time                    `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
