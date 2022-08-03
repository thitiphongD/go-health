package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/go-health/api/orm"
)

func ReadAllAlbum(c *gin.Context) {
	var albums []orm.Album
	orm.Db.Find(&albums)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Read Albums Success",
		"users":   albums,
	})
}
