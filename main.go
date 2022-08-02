package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

		var userExit User

		db.Where("user = ?", json.User).First(&userExit)
		if userExit.ID > 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "User exist",
			})
			return
		}

		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Pass), 10)

		user := User{
			User: json.User,
			Pass: string(encryptedPassword),
			Name: json.Name,
		}

		db.Create(&user)
		if user.ID > 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  "Ok",
				"message": "create success",
				"userId":  user.ID,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "create fail",
			})
		}
	})

	router.Use(cors.Default())

	router.Run()
}
