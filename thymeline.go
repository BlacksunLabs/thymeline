// Thymeline [[ finish this ]]
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func startRouter() {
	router := gin.Default()

	logGroup := router.Group("/log")
	{
		logGroup.POST("/terminal", termLogHandler)
	}

	eventGroup := router.Group("/event")
	{
		eventGroup.POST("/screenshot", scrotEventHandler)
	}

	router.Run(":8080")

}

func termLogHandler(c *gin.Context) {
	// TODO: real-time terminal log handling
	c.JSON(http.StatusOK, gin.H{"Ok": true})
}

func scrotEventHandler(c *gin.Context) {
	// TODO: handle screenshot events
	c.JSON(http.StatusOK, gin.H{"Ok": true})
}
