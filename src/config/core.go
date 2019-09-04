package config

import (
	"fmt"
	"os"
)

// Database - Informations about database connection
type Database struct {
	Name       string
	User       string
	Password   string
	Host       string
	Connection string
	Port       string
}

// Config - General config
type Config struct {
	Address   string
	SecretKey string
	Database  Database
}

var config Config

// LoadConfig - Load all configs from .env file
func LoadConfig() {
	// database credentials
	config.Database.Connection = os.Getenv("DB_CONNECTION")
	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.Name = os.Getenv("DB_DATABASE")
	config.Database.Password = os.Getenv("DB_PASSWORD")
	config.Database.User = os.Getenv("DB_USER")
	// address
	config.Address = fmt.Sprintf(":%s", os.Getenv("PORT"))
	config.SecretKey = os.Getenv("API_KEY")
}

// GetConfig -
func GetConfig() Config {
	return config
}
