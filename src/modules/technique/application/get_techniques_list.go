package application_technique

import domain_technique "jiu-tracker/src/modules/technique/domain"

type GetTechniquesListUseCase struct {
	Repo domain_technique.TechniqueRepository
}

func NewGetTechniquesListUseCase(repo domain_technique.TechniqueRepository) *GetTechniquesListUseCase {
	return &GetTechniquesListUseCase{Repo: repo}
}

func (uc *GetTechniquesListUseCase) Execute() ([]domain_technique.TechniqueListEntry, error) {
	return uc.Repo.FindAllList()
}
