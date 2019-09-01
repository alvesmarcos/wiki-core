package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"wikilibras-core/src/app/models"
)

func connect() (*gorm.DB, error) {
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
func RunMigrations() *gorm.DB {
	db, err := connect()

	if err != nil {
		log.Fatal(err)
		return nil
	}
	db.LogMode(true)

	// migrate models
	db.AutoMigrate(
		&models.User{},
		&models.Context{},
		&models.Action{},
		&models.State{},
		&models.TaskType{},
		&models.Workflow{},
		&models.ObjectType{},
		&models.Object{},
		&models.Task{},
		&models.Assignment{},
		&models.ActionAssignment{},
		&models.Token{},
	)
	// add constraints if model has
	models.AddWorkflowConstraints(db)
	models.AddObjectConstraints(db)
	models.AddTaskConstraints(db)
	models.AddAssignmentConstraints(db)
	models.AddActionAssignmentConstraints(db)
	models.AddTokenConstraints(db)

	return db
}
