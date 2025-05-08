package main

import (
	"bjj-tracker/config"
	domain_user "bjj-tracker/src/modules/user/domain"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&domain_user.User{})
}
