package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	AuthController "github.com/thitiphongD/go-health/api/controllers/auth"
	"github.com/thitiphongD/go-health/orm"
)

func main() {

	orm.InitDB()
	router := gin.Default()
	router.POST("/register", AuthController.Register)
	router.POST("/login", AuthController.Login)
	router.Use(cors.Default())
	router.Run()
}
