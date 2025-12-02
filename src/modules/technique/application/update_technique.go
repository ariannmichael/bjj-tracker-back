package application_technique

import (
	domain_technique "bjj-tracker/src/modules/technique/domain"
	"fmt"
)

type UpdateTechniqueUseCase struct {
	Repo domain_technique.TechniqueRepository
}

func NewUpdateTechniqueUseCase(repo domain_technique.TechniqueRepository) *UpdateTechniqueUseCase {
	return &UpdateTechniqueUseCase{Repo: repo}
}

func (uc *UpdateTechniqueUseCase) Execute(id string, req UpdateTechniqueRequest) (*domain_technique.Technique, error) {
	technique, err := uc.Repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find technique: %w", err)
	}
	technique.Name = req.Name
	technique.NamePortuguese = req.NamePortuguese
	technique.Description = req.Description
	technique.DescriptionPortuguese = req.DescriptionPortuguese
	technique.Category = req.Category
	technique.Difficulty = req.Difficulty
	newTechnique, err := uc.Repo.Update(technique)
	if err != nil {
		return nil, fmt.Errorf("failed to update technique: %w", err)
	}
	if newTechnique == nil {
		return nil, fmt.Errorf("failed to update technique: technique is nil")
	}
	return newTechnique, nil
}
