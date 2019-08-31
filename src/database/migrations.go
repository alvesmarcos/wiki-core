package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"wikilibras-core/src/app/models"
)

func connect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// database credentials
	dialect := os.Getenv("DB_CONNECTION")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	password := os.Getenv("DB_PASSWORD")
	username := os.Getenv("DB_USER")

	return gorm.Open(
		dialect,
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			host, port, username, database, password,
		),
	)
}

// RunMigrations - export tables without erase
func RunMigrations(withSeeds bool) {
	db, err := connect()

	if err != nil {
		log.Fatal(err)

		return
	}
	db.LogMode(true)

	// migrate models
	db.AutoMigrate(
		&models.Action{}, &models.State{}, &models.TaskType{}, &models.Workflow{},
	)
	// add constraints if model has
	models.AddWorkflowConstraints(db)

	// populate tables
	if withSeeds {
		RunSeeds(db)
	}
}
