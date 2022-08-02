package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusOK, gin.H{
			"status":  "Ok",
			"message": "login success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "login fail",
		})
	}
}
