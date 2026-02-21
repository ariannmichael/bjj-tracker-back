package application_training

type CreateTrainingRequest struct {
	UserID                 string   `json:"user_id" binding:"required"`
	Date                   string   `json:"date" binding:"required"`
	IsOpenMat              bool     `json:"is_open_mat" binding:"required"`
	SubmitUsingOptionsIDs  []string `json:"submit_using_options_ids" binding:"required,min=1"`
	SubmittedByOptionsIDs  []string `json:"submitted_by_options_ids" binding:"required,min=1"`
	Duration               int      `json:"duration" binding:"required"` // in minutes
	Notes                  string   `json:"notes"`
}

type UpdateTrainingRequest struct {
	Date                   string   `json:"date" binding:"required"`
	IsOpenMat              bool     `json:"is_open_mat" binding:"required"`
	SubmitUsingOptionsIDs  []string `json:"submit_using_options_ids" binding:"required,min=1"`
	SubmittedByOptionsIDs  []string `json:"submitted_by_options_ids" binding:"required,min=1"`
	Duration               int      `json:"duration" binding:"required"` // in minutes
	Notes                  string   `json:"notes"`
}
