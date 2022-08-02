package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Register struct {
	User string `json:"user" binding:"required"`
	Pass string `json:"pass" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/register", func(c *gin.Context) {
		var json Register
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"register": json,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
