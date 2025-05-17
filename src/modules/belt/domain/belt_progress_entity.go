package domain_belt

import "time"

type Belt int

const (
	White Belt = iota
	Blue
	Purple
	Brown
	Black
)

type BeltProgress struct {
	ID          int       `json:"id" gorm:"primaryKey, autoIncrement"`
	UserID      string    `json:"user_id" gorm:"not null"`
	User        string    `json:"user" gorm:"not null"`
	CurrentBelt Belt      `json:"current_belt" gorm:"not null"`
	StripeCount int       `json:"stripe_count" gorm:"default:0"`
	EarnedAt    time.Time `json:"earned_at" gorm:"default:CURRENT_TIMESTAMP"`
}
