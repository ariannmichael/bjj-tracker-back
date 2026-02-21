package main

import (
	"jiu-tracker/config"
	domain_technique "jiu-tracker/src/modules/technique/domain"
	infrastructure_technique "jiu-tracker/src/modules/technique/infrastructure"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	db := config.DB
	fixtures := infrastructure_technique.TechniqueFixtures()
	log.Printf("Seeding %d techniques...", len(fixtures))

	var created, skipped int
	for _, f := range fixtures {
		var existing domain_technique.Technique
		err := db.Where("id = ?", f.ID).First(&existing).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := db.Create(&f).Error; err != nil {
					log.Printf("Failed to create technique %s: %v", f.ID, err)
					continue
				}
				created++
			} else {
				log.Printf("Error checking technique %s: %v", f.ID, err)
			}
			continue
		}
		skipped++
	}

	log.Printf("Done: %d created, %d skipped (already existed).", created, skipped)
	fmt.Printf("Done: %d created, %d skipped (already existed).\n", created, skipped)
}
