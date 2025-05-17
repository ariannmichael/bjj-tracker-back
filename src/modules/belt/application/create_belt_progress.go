package application_belt

import (
	domain_belt "bjj-tracker/src/modules/belt/domain"
)

type CreateBeltProgressUseCase struct {
	Repo        domain_belt.BeltProgressRepository
	BeltService BeltService
}

func (cu *CreateBeltProgressUseCase) Execute(cbDTO CreateBeltProgressDTO) (*domain_belt.BeltProgress, error) {
	return cu.BeltService.CreateBeltProgress(cbDTO)
}
