package config

import (
	"agmc/day2/internal/domains"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Username     string
	Password     string
	Port         string
	DatabaseName string
}

func CreateConnection(dbConfig DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(db:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Port, dbConfig.DatabaseName)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return DB, err
}

func NewDatabaseConfig(dbConfig DatabaseConfig) *gorm.DB {
	DB, err := CreateConnection(dbConfig)

	if err != nil {
		log.Fatal(err)
	}

	DB.Debug().AutoMigrate(
		&domains.User{},
	)

	return DB
}
