package main

import (
	"jiu-tracker/config"
	domain_belt "jiu-tracker/src/modules/belt/domain"
	domain_technique "jiu-tracker/src/modules/technique/domain"
	domain_training "jiu-tracker/src/modules/training/domain"
	domain_user "jiu-tracker/src/modules/user/domain"
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
