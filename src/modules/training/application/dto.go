package application_training

type CreateTrainingRequest struct {
	UserID       string   `json:"user_id" binding:"required"`
	TechniqueIDs []string `json:"techniques_ids" binding:"required,min=1"`
	Duration     int      `json:"duration" binding:"required"` // in minutes
	Notes        string   `json:"notes"`
}

type UpdateTrainingRequest struct {
	TechniqueIDs []string `json:"techniques_ids" binding:"required,min=1"`
	Duration     int      `json:"duration" binding:"required"` // in minutes
	Notes        string   `json:"notes"`
}
