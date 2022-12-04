package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func API() {
	router := gin.Default()
	router.GET("/search", func(c *gin.Context) {
		query := c.Query("query")
		limit := c.Query("limit")

		c.JSON(http.StatusOK, gin.H{
			"hello": query,
			"limit": limit,
		})
	})
	router.Run(":8080")
}
