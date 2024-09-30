package main

import (
	"github.com/gin-gonic/gin"
)

// All The Application Configurations
type Application struct {
	config Config
}

// Server Configurations
type Config struct {
	addr string
}

func (app *Application) StartServer() {
	// Create New Router
	router := gin.Default()

	// Create /api/v1 Group
	v1 := router.Group("/api/v1")

	// Create /api/v1/ping Route
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	// Start Server
	router.Run(app.config.addr)
}
