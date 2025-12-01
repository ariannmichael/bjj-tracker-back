package application_training

type CreateTrainingRequest struct {
	UserID       string   `json:"user_id" binding:"required"`
	TechniqueIDs []string `json:"technique_ids" binding:"required"`
	Duration     int      `json:"duration" binding:"required"` // in minutes
	Notes        string   `json:"notes"`
}

