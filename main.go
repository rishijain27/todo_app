package main

import (
	"context"
	// "log"
	"os"
	"time"
	"todo_app/config"
	"todo_app/routes"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func main() {
	// loads .env in root directory
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}


	client := config.ConnectDB()

	router := gin.Default()
	router.Use(gin.Logger())
	// CORS middleware
	// i'm not sure this will be necessary but it is here for testing purposes
	// router.Use(middleware.Cors())

	// basic security headers

	routes.SetupRouter(router, client)

	router.Run("localhost:" + port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)

}
