package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ts3486/pet-insurance-api/internal/database"
)

func main() {
	// Load .env file
	godotenv.Load()

	e := echo.New()

	// Database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("database connected")

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	// OpenAPI spec + Swagger UI
	e.GET("/api/v1/openapi.yaml", func(c echo.Context) error {
		return c.File("api/openapi.yaml")
	})
	e.GET("/swagger/", func(c echo.Context) error {
		return c.File("api/swagger.html")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
