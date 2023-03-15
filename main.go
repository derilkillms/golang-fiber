package main

import (
	"golang-fiber/database"
	"golang-fiber/database/migration"
	"golang-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//INITIAL DATABASE
	database.DatabaseInit()
	migration.RunMigration()

	// INITIAL ROUTE
	route.RouteInit(app)
	app.Listen(":3000")
}
