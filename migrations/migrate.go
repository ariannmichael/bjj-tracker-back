package main

import (
	"bjj-tracker/config"
	domain_belt "bjj-tracker/src/modules/belt/domain"
	domain_technique "bjj-tracker/src/modules/technique/domain"
	domain_training "bjj-tracker/src/modules/training/domain"
	domain_user "bjj-tracker/src/modules/user/domain"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&domain_user.User{})
	config.DB.AutoMigrate(&domain_belt.BeltProgress{})
	config.DB.AutoMigrate(&domain_technique.Technique{})
	config.DB.AutoMigrate(&domain_training.TrainingSession{})
}
