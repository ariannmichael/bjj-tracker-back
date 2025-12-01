package domain_training

type TrainingRepository interface {
	CreateTrainingSession(trainingSession *TrainingSession) (*TrainingSession, error)
	GetTrainingSessionByID(id string) (*TrainingSession, error)
	GetAllTrainingSessions() ([]TrainingSession, error)
	UpdateTrainingSession(trainingSession *TrainingSession) (*TrainingSession, error)
	DeleteTrainingSession(id string) error
}
