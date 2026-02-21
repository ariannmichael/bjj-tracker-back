package domain_training

import (
	domain_technique "jiu-tracker/src/modules/technique/domain"
	"time"
)

type TrainingSession struct {
	ID                  string                       `json:"id" gorm:"primaryKey"`
	UserID              string                       `json:"user_id" gorm:"not null"`
	Date                string                       `json:"date" gorm:"not null"`
	IsOpenMat           bool                         `json:"is_open_mat" gorm:"not null"`
	SubmitUsingOptions  []domain_technique.Technique `json:"submit_using_options" gorm:"many2many:training_session_submit_using_options"`
	SubmittedByOptions  []domain_technique.Technique `json:"submitted_by_options" gorm:"many2many:training_session_submitted_by_options"`
	Duration            int                          `json:"duration" gorm:"not null"` // in minutes
	Notes               string                       `json:"notes" gorm:"not null"`
	CreatedAt           time.Time                    `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time                    `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
