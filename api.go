package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func API() {
	router := gin.Default()
	var index map[string][]int

	router.GET("/build", func(c *gin.Context) {
		filename := c.Query("file")
		filename = fmt.Sprintf("%s.json", filename)

		index = BuildIndex(filename)

		c.JSON(200, gin.H{
			"index": "success",
		})
	})

	router.GET("/search", func(c *gin.Context) {
		if len(index) == 0 {
			c.JSON(404, gin.H{
				"Error": "You need to build the index first",
			})
			return
		}

		filename := c.DefaultQuery("file", "nytimes")
		query := c.Query("query")
		limitStr := c.Query("limit")

		limit, _ := strconv.Atoi(limitStr)
		filename = fmt.Sprintf("%s.json", filename)

		results := Search(query, limit, index, filename)
		c.JSON(200, results)
	})

	router.Run(":8080")
}
