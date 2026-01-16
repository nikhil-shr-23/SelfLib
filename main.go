package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello sir",
			"status":"ye bhi okay hai shayad",
		})
	})

	router.Run(":8088")
}
