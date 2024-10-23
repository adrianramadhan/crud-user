package main

import (
	"basic/api/auth"
	"basic/api/handler"
	"basic/api/middleware"
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

	jwtService := auth.NewService()

	// Initialize Gin router
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/login", userHandler.Login(jwtService))

	// api.POST("/users", userHandler.CreateUser)
	// api.GET("/users", userHandler.GetAllUsers)
	// api.GET("/users/:id", userHandler.GetUserByID)
	// api.PUT("/users/:id", userHandler.UpdateUser)
	// api.DELETE("/users/:id", userHandler.DeleteUser)

	// Implement Auth Middleware
	authMiddleware := api.Group("/users")
	authMiddleware.Use(middleware.JWTAuthMiddleware(jwtService))
	{
		authMiddleware.POST("/", userHandler.CreateUser)
		authMiddleware.GET("/", userHandler.GetAllUsers)
		authMiddleware.GET("/:id", userHandler.GetUserByID)
		authMiddleware.PUT("/:id", userHandler.UpdateUser)
		authMiddleware.DELETE("/:id", userHandler.DeleteUser)
	}

	// Start server
	router.Run(":8080")
}
