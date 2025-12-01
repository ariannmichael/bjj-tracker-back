package application_technique

type CreateTechniqueRequest struct {
	Name                  string `json:"name" binding:"required"`
	NamePortuguese        string `json:"name_portuguese" binding:"required"`
	Description           string `json:"description" binding:"required"`
	DescriptionPortuguese string `json:"description_portuguese" binding:"required"`
	Category              string `json:"category" binding:"required"`
	Difficulty            int    `json:"difficulty" binding:"required"`
}

type UpdateTechniqueRequest struct {
	ID                    string `json:"id" binding:"required"`
	Name                  string `json:"name" binding:"required"`
	NamePortuguese        string `json:"name_portuguese" binding:"required"`
	Description           string `json:"description" binding:"required"`
	DescriptionPortuguese string `json:"description_portuguese" binding:"required"`
	Category              string `json:"category" binding:"required"`
	Difficulty            int    `json:"difficulty" binding:"required"`
}
