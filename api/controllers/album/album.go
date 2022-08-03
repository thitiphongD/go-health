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

func CreateAlbum(c *gin.Context) {
	var album orm.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := orm.Db.Create(&album)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Create success",
		"album":   album,
		"result":  result.RowsAffected,
	})
}

// soft delete
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	var deleteAlbum orm.Album
	orm.Db.First(&deleteAlbum, id)
	orm.Db.Delete(&deleteAlbum, id)
	c.JSON(200, deleteAlbum)
}

func UpdateAlbum(c *gin.Context) {
	var album orm.Album
	var updateAlbum orm.Album
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	orm.Db.First(&updateAlbum, album.ID)
	updateAlbum.Imagename = album.Imagename
	updateAlbum.Image = album.Image
	orm.Db.Save(updateAlbum)
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Update success",
		"update":  updateAlbum,
	})
}
