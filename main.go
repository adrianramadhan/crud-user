package main

import (
	"basic/api/handler"
	"basic/api/repository"
	"basic/api/service"
	"basic/config"
	"basic/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to the database!")
	}

	// Migrate the schema (create/update table)
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migration completed!")

	// Initialize repositories, services, and handlers
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Initialize Gin router
	r := gin.Default()

	// Endpoint Create User
	r.POST("/users", userHandler.CreateUser)

	// Endpoint Get All Users
	r.GET("/users", userHandler.GetAllUsers)

	// Endpoint Get User By ID
	r.GET("/users/:id", userHandler.GetUserByID)

	// Start server
	r.Run(":8080")
}
