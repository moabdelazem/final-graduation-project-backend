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
	db   dbConfig
}

// Database base Configurations
type dbConfig struct {
	addr string
}

func (app *Application) Mount() *gin.Engine {
	// Create New Router
	router := gin.Default()

	// Create /api/v1 Group
	v1 := router.Group("/api/v1")

	// Use The Base Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Create /api/v1/ping Route
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	return router
}

func (app *Application) StartServer(router *gin.Engine) {
	// Start Server
	router.Run(app.config.addr)
}
