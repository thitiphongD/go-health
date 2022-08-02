package orm

import "gorm.io/gorm"

type User struct {
	gorm.Model
	User string
	Pass string
	Name string
}
