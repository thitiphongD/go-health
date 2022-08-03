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
		"albums":  albums,
	})
}

func ReadOneAlbum(c *gin.Context) {
	id := c.Param("id")
	var album orm.Album
	orm.Db.First(&album, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Read One Album Success",
		"album":   album,
	})
}
