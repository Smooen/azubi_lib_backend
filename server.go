package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
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

func (h *Handler) getBooks(c echo.Context) error {
	var books []Book

	db := h.DB
	result := db.Find(&books)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Query failed")
	}

	return c.JSON(http.StatusOK, books);
}

func main() {
	err := godotenv.Load("db.env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	e := echo.New()
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	os.Remove("./test.sqlite")
	db, err := gorm.Open(sqlite.Open("./test.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	h := Handler{DB: db} //why even bother with handler if no session?

	migrate(db) // should migrate return err since it might fail?

	var book Book
	db.First(&book)
	log.Printf("Found book: %s\n", book.Title)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	e.GET("/Books", h.getBooks)

	e.Logger.Fatal(e.Start(":1323"))
	
}
