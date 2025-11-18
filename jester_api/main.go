package main 

import (
	"log"
	"github.com/gin-gonic/gin"

	"jester_api/handlers"
	"jester_api/services"
)

func main() {

	//DI - Service
	pingService := services.NewPingService()

	entriesService := services.NewEntriesService()

	//DI - Handler
	pingHandler := handlers.NewPingHandler(pingService)

	entriesHandler := handlers.NewEntriesHandler(entriesService)

	//Webserver
	router := gin.Default();

	router.GET("/ping", pingHandler.Ping)

	router.GET("/entries-current", entriesHandler.GetCurrentEntries)

	// TODO: make the port configurable
	if err := router.Run(":8080"); err != nil { 
		log.Fatalf("server failed %v", err)
	}
	
}