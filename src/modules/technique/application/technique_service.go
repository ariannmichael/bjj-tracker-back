package application_technique

import domain_technique "bjj-tracker/src/modules/technique/domain"

type TechniqueService struct {
	Repo domain_technique.TechniqueRepository
}

func NewTechniqueService(repo domain_technique.TechniqueRepository) *TechniqueService {
	return &TechniqueService{
		Repo: repo,
	}
}

func (ts *TechniqueService) CreateTechnique(techniqueDTO *CreateTechniqueRequest) (*domain_technique.Technique, error) {
	technique := domain_technique.Technique{
		Name:                  techniqueDTO.Name,
		NamePortuguese:        techniqueDTO.NamePortuguese,
		Description:           techniqueDTO.Description,
		DescriptionPortuguese: techniqueDTO.DescriptionPortuguese,
		Category:              techniqueDTO.Category,
		Difficulty:            domain_technique.Difficulty(techniqueDTO.Difficulty),
	}
	return ts.Repo.Create(&technique)
}

func (ts *TechniqueService) UpdateTechnique(techniqueDTO *UpdateTechniqueRequest) (*domain_technique.Technique, error) {
	technique := domain_technique.Technique{
		Name:                  techniqueDTO.Name,
		NamePortuguese:        techniqueDTO.NamePortuguese,
		Description:           techniqueDTO.Description,
		DescriptionPortuguese: techniqueDTO.DescriptionPortuguese,
		Category:              techniqueDTO.Category,
		Difficulty:            domain_technique.Difficulty(techniqueDTO.Difficulty),
	}
	return ts.Repo.Update(&technique)
}

func (ts *TechniqueService) GetTechniqueByID(id string) (*domain_technique.Technique, error) {
	return ts.Repo.FindByID(id)
}

func (ts *TechniqueService) GetTechniquesByCategory(category domain_technique.Category) ([]*domain_technique.Technique, error) {
	techniques, err := ts.Repo.FindByCategory(category)
	if err != nil {
		return nil, err
	}
	techniquePtrs := make([]*domain_technique.Technique, len(techniques))
	for i := range techniques {
		techniquePtrs[i] = &techniques[i]
	}
	return techniquePtrs, nil
}

func (ts *TechniqueService) GetTechniquesByDifficulty(difficulty domain_technique.Difficulty) ([]*domain_technique.Technique, error) {
	techniques, err := ts.Repo.FindByDifficulty(difficulty)
	if err != nil {
		return nil, err
	}
	techniquePtrs := make([]*domain_technique.Technique, len(techniques))
	for i := range techniques {
		techniquePtrs[i] = &techniques[i]
	}
	return techniquePtrs, nil
}

func (ts *TechniqueService) GetTechniquesByIDs(ids []string) ([]domain_technique.Technique, error) {
	return ts.Repo.FindByIDs(ids)
}

func (ts *TechniqueService) GetAllTechniques() ([]*domain_technique.Technique, error) {
	techniques, err := ts.Repo.FindAll()
	if err != nil {
		return nil, err
	}
	techniquePtrs := make([]*domain_technique.Technique, len(techniques))
	for i := range techniques {
		techniquePtrs[i] = &techniques[i]
	}
	return techniquePtrs, nil
}
