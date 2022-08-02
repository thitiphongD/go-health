package user

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/thitiphongD/go-health/orm"
)

var HealthHomeSecret []byte

func ReadAllUser(c *gin.Context) {
	HealthHomeSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	header := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(header, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return HealthHomeSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var users []orm.User
		fmt.Println(claims["userId"])
		orm.Db.Find(&users)
		c.JSON(http.StatusOK, gin.H{
			"status":  "Ok",
			"message": "read success",
			"user":    users,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "forbidden",
			"message": err.Error(),
		})
		return
	}

}
