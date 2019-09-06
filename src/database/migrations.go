package database

import (
	"fmt"
	"log"
	"wikilibras-core/src/config"

	"github.com/jinzhu/gorm"
	// dialect gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"wikilibras-core/src/app/models"
)

// RunMigrations - export tables without erase
func RunMigrations() *gorm.DB {
	db, err := connect()

	if err != nil {
		log.Fatal(err)
		return nil
	}
	db.LogMode(true)

	// WARNING: this will drop tables without store your data
	// drop tables
	dropTables(db)

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
		&models.ActionState{},
		&models.Token{},
		&models.Orientation{},
		&models.Profile{},
		&models.UserProfile{},
	)
	// add constraints if model has
	models.AddWorkflowConstraints(db)
	models.AddObjectConstraints(db)
	models.AddTaskConstraints(db)
	models.AddAssignmentConstraints(db)
	models.AddActionStateConstraints(db)
	models.AddTokenConstraints(db)
	models.AddUserProfileConstraints(db)

	return db
}

func connect() (*gorm.DB, error) {
	return gorm.Open(
		config.GetConfig().Database.Connection,
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.GetConfig().Database.Host,
			config.GetConfig().Database.Port,
			config.GetConfig().Database.User,
			config.GetConfig().Database.Name,
			config.GetConfig().Database.Password,
		),
	)
}

func dropTables(db *gorm.DB) {
	db.DropTableIfExists(
		&models.Token{},
		&models.ActionState{},
		&models.Assignment{},
		&models.Object{},
		&models.ObjectType{},
		&models.Workflow{},
		&models.TaskType{},
		&models.Task{},
		&models.State{},
		&models.Action{},
		&models.Context{},
		&models.UserProfile{},
		&models.User{},
		&models.Orientation{},
		&models.Profile{},
	)
}
