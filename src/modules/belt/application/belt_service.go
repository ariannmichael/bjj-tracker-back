package application_belt

import (
	domain_belt "bjj-tracker/src/modules/belt/domain"
	"fmt"
	"strings"
)

type BeltService struct {
	Repo domain_belt.BeltProgressRepository
}

func NewBeltService(repo domain_belt.BeltProgressRepository) *BeltService {
	return &BeltService{
		Repo: repo,
	}
}

func (b *BeltService) GetBeltByColor(color string) (domain_belt.Belt, error) {
	var belt domain_belt.Belt = domain_belt.White

	switch strings.ToLower(color) {
	case "white":
		belt = domain_belt.White
	case "blue":
		belt = domain_belt.Blue
	case "purple":
		belt = domain_belt.Purple
	case "brown":
		belt = domain_belt.Brown
	case "black":
		belt = domain_belt.Black
	default:
		return 0, fmt.Errorf("invalid belt color: %s", color)
	}

	return belt, nil
}

func (b *BeltService) CreateBeltProgress(cbDTO CreateBeltProgressDTO) (*domain_belt.BeltProgress, error) {
	belt, err := b.GetBeltByColor(cbDTO.Color)

	if err != nil {
		return nil, err
	}

	beltProgress := domain_belt.BeltProgress{
		UserID:      cbDTO.UserID,
		CurrentBelt: belt,
		StripeCount: cbDTO.Stripes,
	}

	createdBeltProgress, err := b.Repo.CreateBeltProgress(&beltProgress)
	if err != nil {
		return nil, err
	}
	return createdBeltProgress, nil
}
