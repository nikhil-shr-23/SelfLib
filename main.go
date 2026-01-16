package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello sir",
			"status":"ye bhi okay hai shayad"
		})
	})

	r.Run(":8088")
}
