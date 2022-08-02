package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Register struct {
	User string `json:"user" binding:"required"`
	Pass string `json:"pass" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type User struct {
	gorm.Model
	User string
	Pass string
	Name string
}

func main() {
	router := gin.Default()

	dsn := "root:daew@tcp(127.0.0.1:3306)/go_health?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

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

	router.Run()
}
