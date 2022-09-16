package main

import (
	"agmc/day2/config"
	"agmc/day2/middleware"
	"agmc/day2/routes"
)

func main() {
	DB := config.DatabaseConfig.GetDatabaseConfig(config.DatabaseConfig{})

	defer DB.DB()

	pub, priv, err := middleware.GetRSAKeys()

	if err != nil {
		panic(err)
	}

	r := routes.NewRouting()

	e := r.GetRouting(priv, pub)

	e.Logger.Fatal(e.Start(":1323"))
}
