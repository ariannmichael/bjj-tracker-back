package application_technique

import (
	"bjj-tracker/config"
	domain_technique "bjj-tracker/src/modules/technique/domain"
	infrastructure_technique "bjj-tracker/src/modules/technique/infrastructure"
)

type CreateTechniqueUseCase struct {
	Repo             domain_technique.TechniqueRepository
	TechniqueService *TechniqueService
}

func NewCreateTechniqueUseCase() *CreateTechniqueUseCase {
	db := config.ConnectToDB()
	repo := &infrastructure_technique.TechniqueRepositoryImpl{DB: db}
	techniqueService := NewTechniqueService(repo)

	return &CreateTechniqueUseCase{
		Repo:             repo,
		TechniqueService: techniqueService,
	}
}

func (uc *CreateTechniqueUseCase) Execute(req CreateTechniqueRequest) (*domain_technique.Technique, error) {
	technique := domain_technique.Technique{
		Name:                  req.Name,
		NamePortuguese:        req.NamePortuguese,
		Description:           req.Description,
		DescriptionPortuguese: req.DescriptionPortuguese,
		Category:              req.Category,
		Difficulty:            domain_technique.Difficulty(req.Difficulty),
	}
	return uc.Repo.Create(&technique)
}
