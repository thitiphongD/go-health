package orm

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Id        uint
	Imagename string
	Image     string
}
