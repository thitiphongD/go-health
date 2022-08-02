package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/go-health/orm"
)

func ReadAllUser(c *gin.Context) {
	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Read User Success",
		"users":   users,
	})
}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user orm.User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Read Profile Success",
		"user":    user,
	})
}
