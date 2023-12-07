package main

import (
	"github.com/franBarrientos/infrastructure/config"
	"github.com/franBarrientos/infrastructure/gorm"
	"github.com/franBarrientos/infrastructure/rest/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	env := config.GetEnv()
	db, err := gorm.InitDbConnection(env)

	if err != nil {
		log.Fatal("Failed to initialize database connection:", err)
		return
	}

	app := fiber.New()
	routes.InitRoutes(app, db)
	app.Listen(":3000")
}
