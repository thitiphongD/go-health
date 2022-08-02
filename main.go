package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	AuthController "github.com/thitiphongD/go-health/api/controllers/auth"
	UserController "github.com/thitiphongD/go-health/api/controllers/user"
	"github.com/thitiphongD/go-health/orm"
)

func main() {
	router := gin.Default()
	orm.InitDB()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.POST("/register", AuthController.Register)
	router.POST("/login", AuthController.Login)
	router.GET("/users/read-all", UserController.ReadAllUser)
	router.Use(cors.Default())
	router.Run()
}
