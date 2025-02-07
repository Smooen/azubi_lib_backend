package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique" json:"username" query:"username"`
	Password string `gorm:"type:varchar(100)" json:"-"`
	Email    string `gorm:"type:varchar(100);unique" json:"email" query:"email"`
	Favorites Favorites
}

type Book struct {
	gorm.Model
	Title  string		`json:"title" query:"title"`
	Isbn string			`json:"isbn" query:"isbn"`
	Author string		`json:"author" query:"author"`
	ReleaseDate string	`json:"releaseDate" query:"releaseDate"`
	Availability bool	`json:"availability" query:"availability"`
}

type Favorites struct {
	gorm.Model
	UserID uint
	Books []Book `gorm:"many2many:book_favorites"`
}
