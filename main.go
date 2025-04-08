package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"goApiTask/database"
	"goApiTask/database/query"
	"goApiTask/handlers"
	"goApiTask/routes"
)

func main() {
	app := fiber.New()

	conn, err := database.Connect()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	repo := query.NewRepository(conn)
	handl := handlers.NewClient(repo)

	routes.RegisterTaskRoutes(app, handl)

	log.Fatal(app.Listen(":8080"))
	defer conn.Close(context.Background())
}
