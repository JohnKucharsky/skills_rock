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
	"syscall"
)

// @title Tasks
// @version 1.0
// @description These are docs for test assignment from Skills Rock.
// @host localhost:8080
// @BasePath /api

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Can't load env")
	}

	// Get Postgres URI from .env
	postgresURI := os.Getenv("POSTGRES_URI")
	if postgresURI == "" {
		log.Fatal("Error: POSTGRES_URI is not set")
	}

	// Initialize Fiber app
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

	// Serve Swagger docs
	app.Get("/docs/*", swagger.HandlerDefault)

	// Connect to the database
	database, err := db.New(postgresURI)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Register internal routes
	internal.Register(app, database)

	// Graceful shutdown handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nGracefully shutting down...")
		_ = app.Shutdown()
	}()

	// Start the server
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
