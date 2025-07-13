package main

import (
	"fmt"
	"beTestBookingToGo/pkg/config"
	"beTestBookingToGo/pkg/database"
	httpDelivery "beTestBookingToGo/pkg/delivery/http"
	"beTestBookingToGo/pkg/middleware"
	"log"
	"net/http"

	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// Load Configurations
	config.LoadAppConfig()

	// Initialize Database
	database.Connect(config.AppConfig.ConnectionString)
	database.Migrate()

	// Initialize router
	router := httpDelivery.InitRoutes()
	
	// Wrap router with CORS middleware
	handlerWithCORS := middleware.CORS(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), handlerWithCORS))
}
