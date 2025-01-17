package main

import (
	"azubi_library/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	tx := db.Session(&gorm.Session{})

	tx.AutoMigrate(&models.User{})
	tx.AutoMigrate(&models.Book{})
	tx.AutoMigrate(&models.Favorites{})

	createTestData(tx)
}

func createTestData(tx *gorm.DB) {

	tx.Create(&models.Book{
		Title:  "The Catcher in the Rye",
		Isbn: "9781438119250",
		Author: "J. D. Salinger",
		ReleaseDate: "July 16, 1951",
		Availability: true,
	})

	tx.Create(&models.Book{
		Title:  "The Lord of the Rings",
		Isbn: "9780007322596",
		Author: "J. R. R. Tolkien",
		ReleaseDate: "July 29, 1954",
		Availability: false,
	})
	
	tx.Create(&models.Book{
		Title:  "Pride and Prejudice",
		Isbn: "9780198826736",
		Author: "Jane Austen",
		ReleaseDate: "January 28, 1813",
		Availability: false,
	})
	tx.Create(&models.Book{
		Title:  "1984",
		Isbn: "9783641279110",
		Author: "George Orwell",
		ReleaseDate: "1949",
		Availability: true,
	})

	tx.Create(&models.User{
		Username: "Susanne",
		Password: "1234",
		Email: "test@mail.com",
	})
}
