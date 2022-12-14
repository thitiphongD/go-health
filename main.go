package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	CreateAlbumController "github.com/thitiphongD/go-health/api/controllers/album"
	DeleteAlbumController "github.com/thitiphongD/go-health/api/controllers/album"
	ReadAllAlbumController "github.com/thitiphongD/go-health/api/controllers/album"
	ReadOneAlbumController "github.com/thitiphongD/go-health/api/controllers/album"
	UpdateAlbumController "github.com/thitiphongD/go-health/api/controllers/album"
	AuthController "github.com/thitiphongD/go-health/api/controllers/auth"
	UserController "github.com/thitiphongD/go-health/api/controllers/user"
	"github.com/thitiphongD/go-health/api/middlewares"
	"github.com/thitiphongD/go-health/api/orm"
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
	authorized := router.Group("/users", middlewares.JWTAuthen())
	authorized.GET("/readall", UserController.ReadAllUser)
	authorized.GET("/profile", UserController.Profile)

	router.GET("/albums", ReadAllAlbumController.ReadAllAlbum)
	router.GET("/albums/:id", ReadOneAlbumController.ReadOneAlbum)
	router.POST("/albums", CreateAlbumController.CreateAlbum)
	router.PATCH("/albums", UpdateAlbumController.UpdateAlbum)
	router.DELETE("/albums/:id", DeleteAlbumController.DeleteAlbum)
	router.Use(cors.Default())
	router.Run()
}
