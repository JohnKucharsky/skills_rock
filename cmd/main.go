package main

import (
	"fmt"
	_ "github.com/JohnKucharsky/skills_rock/cmd/docs"
	"github.com/JohnKucharsky/skills_rock/internal"
	"github.com/JohnKucharsky/skills_rock/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

// @title Tasks
// @version 1.0
// @description These are docs for test assignment from Skills Rock.
// @host localhost:8080
// @BasePath /api

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Can't load env")
	}
	postgresURI := os.Getenv("POSTGRES_URI")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(
		cors.New(
			cors.Config{
				AllowOrigins: "*",
				AllowHeaders: "Origin, Content-Type, Accept, Authorization",
				AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
			},
		),
	)

	app.Get("/docs/*", swagger.HandlerDefault)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()
	database := db.New(postgresURI)

	internal.Register(app, database)

	err := app.Listen(":8080")
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
}
