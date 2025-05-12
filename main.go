package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/message", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test OK",
		})
	})

	r.Run() // écoute sur :8080 par défaut
}
