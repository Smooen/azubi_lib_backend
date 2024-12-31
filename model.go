package main

import (
	"gorm.io/gorm"
)

type (
	Handler struct {
		DB *gorm.DB
	}
)

type Book struct {
	gorm.Model
	Title  string		`json:"title" query:"title"`
	Isbn string			`json:"isbn" query:"isbn"`
	Author string		`json:"author" query:"author"`
	ReleaseDate string	`json:"releaseDate" query:"releaseDate"`
	Availability bool	`json:"availability" query:"availability"`
}
