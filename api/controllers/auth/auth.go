package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/thitiphongD/go-health/orm"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	User string `json:"user" binding:"required"`
	Pass string `json:"pass" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var userExit orm.User

	orm.Db.Where("user = ?", json.User).First(&userExit)
	if userExit.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "User exist",
		})
		return
	}

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Pass), 10)

	user := orm.User{
		User: json.User,
		Pass: string(encryptedPassword),
		Name: json.Name,
	}

	orm.Db.Create(&user)
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
}

var HealthHomeSecret []byte

type LoginBody struct {
	User string `json:"user" binding:"required"`
	Pass string `json:"pass" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var userExit orm.User

	orm.Db.Where("user = ?", json.User).First(&userExit)
	if userExit.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "User does not exist",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userExit.Pass), []byte(json.Pass))

	if err == nil {
		HealthHomeSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExit.ID,
			"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(HealthHomeSecret)

		fmt.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{
			"status":  "Ok",
			"message": "login success",
			"token":   tokenString,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "login fail",
		})
	}
}
