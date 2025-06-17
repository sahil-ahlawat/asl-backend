package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/go-sql-driver/mysql"

	"fitness/config"
	"fitness/routes"
)
func main() {
	config.LoadEnv()
	app := fiber.New()

	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/fitness") // Change credentials as needed
	if err != nil {
		log.Fatal(err)
	}

	routes.SetupListingRoutes(app, db)

	port := config.GetEnv("APP_PORT", "3000")
	log.Fatal(app.Listen(":" + port))
}
