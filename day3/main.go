package main

import (
	"agmc/day2/config"
	"agmc/day2/routes"
)

func main() {
	DB := config.DatabaseConfig.GetDatabaseConfig(config.DatabaseConfig{})

	defer DB.DB()

	e := routes.Routing.GetRouting(routes.Routing{})
	e.Logger.Fatal(e.Start(":1323"))
}
