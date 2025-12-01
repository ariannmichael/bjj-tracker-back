package application_technique

import (
	domain_technique "bjj-tracker/src/modules/technique/domain"
)

type GetTechniqueByIDUseCase struct {
	Repo domain_technique.TechniqueRepository
}

func NewGetTechniqueByIDUseCase(repo domain_technique.TechniqueRepository) *GetTechniqueByIDUseCase {
	return &GetTechniqueByIDUseCase{Repo: repo}
}

func (uc *GetTechniqueByIDUseCase) Execute(id string) (*domain_technique.Technique, error) {
	return uc.Repo.FindByID(id)
}
