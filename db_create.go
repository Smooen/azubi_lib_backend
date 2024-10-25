package main

import (
	//"database/sql"

	//"github.com/labstack/gommon/log"

	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {

	db.AutoMigrate(&Book{})

	createTestData(db)
}

func createTestData(db *gorm.DB) {

	db.Create(&Book{
		Title:  "The Catcher in the Rye",
		Isbn: "9781438119250",
		Author: "J. D. Salinger",
		ReleaseDate: "July 16, 1951",
		Availability: true,
	})

	db.Create(&Book{
		Title:  "The Lord of the Rings",
		Isbn: "9780007322596",
		Author: "J. R. R. Tolkien",
		ReleaseDate: "July 29, 1954",
		Availability: false,
	})
	
	db.Create(&Book{
		Title:  "Pride and Prejudice",
		Isbn: "9780198826736",
		Author: "Jane Austen",
		ReleaseDate: "January 28, 1813",
		Availability: false,
	})
	db.Create(&Book{
		Title:  "1984",
		Isbn: "9783641279110",
		Author: "George Orwell",
		ReleaseDate: "1949",
		Availability: true,
	})
}
