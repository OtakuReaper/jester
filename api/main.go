package main 

import (
	"context"
	"log"
	"jester_api/database"

	"jester_api/controllers"
	"github.com/gin-gonic/gin"
	"jester_api/repository"
)

func main() {

	//initialize database
	ctx := context.Background()
	db, err := database.InitDatabase(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	//setting up webserver
	server := gin.Default()

	//endpoints
	server.GET("/users", controllers.GetAllUsers)

	//start server
	server.Run(":3000")
}