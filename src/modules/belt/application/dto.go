package application_belt

type CreateBeltProgressDTO struct {
	UserID  string `json:"user_id" binding:"required"`
	Color   string `json:"color" binding:"required"`
	Stripes int    `json:"stripes" binding:"required"`
}
