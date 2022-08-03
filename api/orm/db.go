package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB() {
	//dsn := os.Getenv("MYSQL_DSN")
	dsn := "root:daew@tcp(127.0.0.1:3306)/go_health?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Album{})

	// Db.Create(&Album{
	// 	Imagename: "Cat",
	// 	Image:     "https://pbs.twimg.com/media/DhxQ4RoWsAAgyww?format=jpg&name=medium",
	// })
}
