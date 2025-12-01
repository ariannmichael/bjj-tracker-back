package domain_technique

type Difficulty int

const (
	Beginner Difficulty = iota
	Intermediate
	Advanced
)

type Category int

const (
	Submission Category = iota
	Sweep
	Pass
	Guard
	Takedown
	Defend
	SubmissionEscape
)

type Technique struct {
	ID                    string     `json:"id" gorm:"primaryKey"`
	Name                  string     `json:"name" gorm:"not null, unique" `
	NamePortuguese        string     `json:"name_portuguese" gorm:"not null"`
	Description           string     `json:"description" gorm:"not null"`
	DescriptionPortuguese string     `json:"description_portuguese" gorm:"not null"`
	Category              string     `json:"category" gorm:"not null"`
	Difficulty            Difficulty `json:"difficulty" gorm:"not null"`
	TrainingSessionID     string     `json:"training_session_id" gorm:"not null"`
	TrainingSession       string     `json:"training_session" gorm:"not null"`
}
