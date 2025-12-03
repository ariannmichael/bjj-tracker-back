package infrastructure_technique

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	domain_technique "bjj-tracker/src/modules/technique/domain"
)

type TechniqueRepositoryImpl struct {
	DB *gorm.DB
}

var _ domain_technique.TechniqueRepository = &TechniqueRepositoryImpl{}

func (r *TechniqueRepositoryImpl) Create(technique *domain_technique.Technique) (*domain_technique.Technique, error) {
	var existingTechnique domain_technique.Technique
	if err := r.DB.Where("name = ?", technique.Name).First(&existingTechnique).Error; err == nil {
		return nil, fmt.Errorf("technique with name %s already exists", technique.Name)
	}

	technique.ID = uuid.New().String()

	if err := r.DB.Create(technique).Error; err != nil {
		return nil, fmt.Errorf("REPO failed to create technique: %w", err)
	}

	return technique, nil
}

func (r *TechniqueRepositoryImpl) Update(technique *domain_technique.Technique) (*domain_technique.Technique, error) {
	if err := r.DB.Save(technique).Error; err != nil {
		return nil, fmt.Errorf("failed to update technique: %w", err)
	}
	return technique, nil
}

func (r *TechniqueRepositoryImpl) FindByID(id string) (*domain_technique.Technique, error) {
	var technique domain_technique.Technique
	if err := r.DB.First(&technique, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("REPO failed to find technique: %w", err)
	}
	return &technique, nil
}

func (r *TechniqueRepositoryImpl) FindByCategory(category domain_technique.Category) ([]domain_technique.Technique, error) {
	var techniques []domain_technique.Technique
	if err := r.DB.Where("category = ?", category).Find(&techniques).Error; err != nil {
		return nil, fmt.Errorf("REPO failed to find techniques: %w", err)
	}
	return techniques, nil
}

func (r *TechniqueRepositoryImpl) FindByDifficulty(difficulty domain_technique.Difficulty) ([]domain_technique.Technique, error) {
	var techniques []domain_technique.Technique
	if err := r.DB.Where("difficulty = ?", difficulty).Find(&techniques).Error; err != nil {
		return nil, fmt.Errorf("REPO failed to find techniques: %w", err)
	}
	return techniques, nil
}

func (r *TechniqueRepositoryImpl) FindByIDs(ids []string) ([]domain_technique.Technique, error) {
	var techniques []domain_technique.Technique
	if err := r.DB.Where("id IN ?", ids).Find(&techniques).Error; err != nil {
		return nil, fmt.Errorf("REPO failed to find techniques: %w", err)
	}
	return techniques, nil
}

func (r *TechniqueRepositoryImpl) FindAll() ([]domain_technique.Technique, error) {
	var techniques []domain_technique.Technique
	if err := r.DB.Find(&techniques).Error; err != nil {
		return nil, fmt.Errorf("REPO failed to find techniques: %w", err)
	}
	return techniques, nil
}
