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
	Host         string
}

func CreateConnection(dbConfig DatabaseConfig) (*gorm.DB, error) {
	var port = ""

	if len(dbConfig.Port) != 0 {
		port = ":" + dbConfig.Port
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?tls=true", dbConfig.Username, dbConfig.Password, dbConfig.Host, port, dbConfig.DatabaseName)
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
