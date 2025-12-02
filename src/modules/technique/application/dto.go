package application_technique

import domain_technique "bjj-tracker/src/modules/technique/domain"

type CreateTechniqueRequest struct {
	Name                  string                      `json:"name" binding:"required"`
	NamePortuguese        string                      `json:"name_portuguese" binding:"required"`
	Description           string                      `json:"description" binding:"required"`
	DescriptionPortuguese string                      `json:"description_portuguese" binding:"required"`
	Category              domain_technique.Category   `json:"category" binding:"required"`
	Difficulty            domain_technique.Difficulty `json:"difficulty" binding:"required"`
}

type UpdateTechniqueRequest struct {
	ID                    string                      `json:"id" binding:"required"`
	Name                  string                      `json:"name" binding:"required"`
	NamePortuguese        string                      `json:"name_portuguese" binding:"required"`
	Description           string                      `json:"description" binding:"required"`
	DescriptionPortuguese string                      `json:"description_portuguese" binding:"required"`
	Category              domain_technique.Category   `json:"category" binding:"required"`
	Difficulty            domain_technique.Difficulty `json:"difficulty" binding:"required"`
}
