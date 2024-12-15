package main

import (
	"dev-mgmt/internal/infrastructure/database"
	"dev-mgmt/internal/infrastructure/server"
	"dev-mgmt/internal/interface/handler"
	"dev-mgmt/internal/usecase"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// MySQL database connection
	db := database.NewMySQLDB()

	// Run database migration
	if err := database.Migrate(db); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	// Initialize repository, usecase, and handler
	repo := database.NewDeviceRepositoryImpl(db)
	usecase := usecase.NewDeviceUseCase(repo)
	handler := handler.NewDeviceHandler(usecase)

	// Setup Gin router
	r := server.NewRouter(handler)

	// Start the server
	r.Run(":8080")
}
