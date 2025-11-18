package main 

import (
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"

	"jester_api/handlers"
	"jester_api/services"
)

func main() {

	//DI - Service
	pingService := services.NewPingService()

	// entriesService := services.NewEntriesService()

	//DI - Handler
	pingHandler := handlers.NewPingHandler(pingService)

	// entriesHandler := handlers.NewEntriesHandler(entriesService)

	//Webserver
	router := gin.Default();

	//CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"}, //allowed domains
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge: 12* time.Hour,
	}))

	router.GET("/ping", pingHandler.Ping)

	// router.GET("/entries-current", entriesHandler.GetCurrentEntries)

	// TODO: make the port configurable
	if err := router.Run(":8080"); err != nil { 
		log.Fatalf("server failed %v", err)
	}
	
}