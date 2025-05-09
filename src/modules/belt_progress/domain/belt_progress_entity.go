package domain_belt

type Belt int

const (
	White Belt = iota
	Blue
	Purple
	Brown
	Black
)

type BeltProgress struct {
	ID          string `json:"id" gorm:"primaryKey"`
	UserID      string `json:"user_id" gorm:"not null"`
	User        string `json:"user" gorm:"not null"`
	CurrentBelt Belt   `json:"current_belt" gorm:"not null"`
	StripeCount int    `json:"stripe_count" gorm:"default:0"`
	EarnedAt    string `json:"earned_at" gorm:"not null"`
}
