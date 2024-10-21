package main

import (
	"basic/api/handler"
	"basic/api/repository"
	"basic/api/service"
	"basic/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to the database!")
	}

	// Initialize repositories, services, and handlers
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Initialize Gin router
	r := gin.Default()

	// Define routes
	r.POST("/users", userHandler.CreateUser)
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	// Start server
	r.Run(":8080")
}
