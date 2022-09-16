package config

import (
	"agmc/day2/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct{}

func (dbConfig DatabaseConfig) GetDatabaseConfig() *gorm.DB {
	dsn := "webuser:webpass@tcp(db:3306)/test_db?charset=utf8&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB.Debug().AutoMigrate(
		&entity.User{},
	)

	return DB
}
