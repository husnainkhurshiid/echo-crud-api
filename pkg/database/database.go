package database

import (
	"employee-manager/internal/model"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Driver           string `yaml:"db_driver"`
	ConnectionString string `yaml:"connection_string"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

// to load db credentials from config.yaml file
func loadConfig(configPath string) (Config, error) {
	var config Config

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// to eastablish a connection with sqlite db
func OpenDbConnection() (*gorm.DB, error) {
	filePath := os.Getenv("CONFIG_PATH")
	_, err := loadConfig(filePath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	connUrl := fmt.Sprintf("file:%s?cache=shared&_loc=auto", os.Getenv("DB_PATH"))
	db, err := gorm.Open(sqlite.Open(connUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.Employee{})
	return db, nil
}

// function to close database connection
func CloseDbConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	sqlDB.Close()
}

// to eastablish a connection with sqlite db
func OpenDbConnectionTest() (*gorm.DB, error) {
	filePath := "../configs/config.yaml"
	_, err := loadConfig(filePath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	connUrl := fmt.Sprintf("file:%s?cache=shared&_loc=auto", os.Getenv("DB_PATH"))
	db, err := gorm.Open(sqlite.Open(connUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.Employee{})
	return db, nil
}
