package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	docs "github.com/roh4nyh/airbnb-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Airbnb API (SaffronStays)
// @version 1.0
// @description This is a Assignment server Airbnb API (SaffronStays).
// @termsOfService http://swagger.io/terms/
// @host airbnb-api-saffronstays.onrender.com
// @BasePath /api/v1
func main() {
	app := gin.Default()
	PORT := "8080"
	app.Use(gin.Logger())
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"https"}

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	app.Use(cors.New(config))

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiV1 := app.Group("/api/v1")
	apiV1.GET("/health", HealthCheckHandler)
	apiV1.GET("/getMetrics/:room_id", GetMetricsHandler)

	log.Printf("server running on port %s", PORT)
	app.Run(fmt.Sprintf(":%s", PORT))
}

// Health check route
// @Summary     Health check
// @Description Health check endpoint to see if the server is running
// @Tags        Health
// @Accept      json
// @Produce     json
// @Success     200 {object} map[string]string
// @Router      /health [get]
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "server is up and running..."})
}

// GetMetricsHandler fetches availability and price data and calculates metrics
// @Summary     Get metrics
// @Description Fetches availability and price data and calculates metrics
// @Tags        Metrics
// @Accept      json
// @Produce     json
// @Param       room_id path string true "Room ID"
// @Success     200 {object} APIResponse
// @Router      /getMetrics/{room_id} [get]
func GetMetricsHandler(c *gin.Context) {
	roomID := c.Param("room_id")
	if roomID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_id is required"})
		return
	}

	availability, err := FetchAvailability(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	prices, err := FetchPrice(roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	metrics := CalculateMetrics(roomID, availability, prices)

	c.JSON(http.StatusOK, metrics)
}
