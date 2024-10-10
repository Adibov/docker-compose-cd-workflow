package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/echo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World!")
	})

	router.Run(":8080")
}
