package application_technique

import (
	domain_technique "bjj-tracker/src/modules/technique/domain"
)

type GetAllTechniquesUseCase struct {
	Repo domain_technique.TechniqueRepository
}

func NewGetAllTechniquesUseCase(repo domain_technique.TechniqueRepository) *GetAllTechniquesUseCase {
	return &GetAllTechniquesUseCase{Repo: repo}
}

func (uc *GetAllTechniquesUseCase) Execute() ([]*domain_technique.Technique, error) {
	techniques, err := uc.Repo.FindAll()
	if err != nil {
		return nil, err
	}
	techniquePtrs := make([]*domain_technique.Technique, len(techniques))
	for i := range techniques {
		techniquePtrs[i] = &techniques[i]
	}
	return techniquePtrs, nil
}
