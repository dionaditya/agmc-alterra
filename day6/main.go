package main

import (
	"agmc/day2/database/config"
	"agmc/day2/internal/factory"
	"agmc/day2/internal/middleware"
	"agmc/day2/internal/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	DB := config.NewDatabaseConfig(config.DatabaseConfig{
		Username:     os.Getenv("DB_USERNAME"),
		Password:     os.Getenv("DB_PASS"),
		Port:         os.Getenv("DB_PORT"),
		DatabaseName: os.Getenv("DB_NAME"),
	})

	defer DB.DB()

	pub, priv, err := middleware.GetRSAKeys()

	if err != nil {
		panic(err)
	}

	f := factory.NewFactory(DB)

	r := routes.NewRouting(f)

	e := r.GetRouting(priv, pub)

	e.Logger.Fatal(e.Start(":1323"))
}
