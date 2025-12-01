package domain_technique

type TechniqueRepository interface {
	Create(technique *Technique) (*Technique, error)
	Update(technique *Technique) (*Technique, error)
	FindByID(id string) (*Technique, error)
	FindByIDs(ids []string) ([]Technique, error)
	FindByCategory(category Category) ([]Technique, error)
	FindByDifficulty(difficulty Difficulty) ([]Technique, error)
	FindAll() ([]Technique, error)
}
