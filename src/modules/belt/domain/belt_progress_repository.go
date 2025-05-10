package domain_belt

type BeltProgressRepository interface {
	CreateBeltProgress(beltProgress *BeltProgress) (*BeltProgress, error)
}
