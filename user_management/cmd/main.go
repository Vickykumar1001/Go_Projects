package main

import (
	"log"
	"user_management/internal/handlers"
	"user_management/internal/middlewares"
	"user_management/internal/repository"
	"user_management/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// We will load the env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot find env file")
		return
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Handler Object
	repo := repository.NewInMemory()
	jwtService := &services.JWTService{}
	handler := handlers.NewHandler(repo, jwtService)
	v1 := r.Group("/api/v1")
	// Two types of groups // auth routes
	auth := v1.Group("/auth") // /api/v1/auth

	auth.POST("/signup", handler.Signup)
	auth.POST("/login", handler.Login)

	// user routes
	user := v1.Group("/user") // /api/v1/user
	user.GET("/getUsers", middlewares.AuthorizationMiddleware(), handler.GetAllUsers)
	user.GET("/get/:username", middlewares.AuthorizationMiddleware(), handler.GetUser)
	user.PUT("/update/:username", middlewares.AuthorizationMiddleware(), handler.UpdateUser)
	user.DELETE("/delete/:username", middlewares.AuthorizationMiddleware(), handler.DeleteUser)

	err = r.Run("localhost:8080")
	if err != nil {
		return
	}

}
