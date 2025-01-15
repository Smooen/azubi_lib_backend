package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (h *Handler) getBooks(c echo.Context) error {
	var books []Book

	id := c.QueryParam("id")
	title := c.QueryParam("title")
	isbn := c.QueryParam("isbn")
	author := c.QueryParam("author")
	releaseDate := c.QueryParam("releaseDate")
	availability := c.QueryParam("availability")

	tx := h.DB.Session(&gorm.Session{})

	if id != "" {
		tx = tx.Where("id = ?", id)
	}
	if title != "" {
		tx = tx.Where("title = ?", title)
	}
	if isbn != "" {
		tx = tx.Where("isbn = ?", isbn)
	}
	if author != "" {
		tx = tx.Where("author = ?", author)
	}
	if releaseDate != "" {
		tx = tx.Where("release_date = ?", releaseDate)
	}
	if availability != "" {
		tx = tx.Where("availability = ?", availability)
	}

	result := tx.Find(&books)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Query failed")
	}

	return c.JSON(http.StatusOK, books)
}

func (h *Handler) getBook(c echo.Context) error {
	var book Book

	id := c.QueryParam("id")
	title := c.QueryParam("title")
	isbn := c.QueryParam("isbn")
	author := c.QueryParam("author")
	releaseDate := c.QueryParam("releaseDate")
	availability := c.QueryParam("availability")

	tx := h.DB.Session(&gorm.Session{})
	query := tx

	if id != "" {
		query = query.Where("id = ?", id)
	}
	if title != "" {
		query = query.Where("title = ?", title)
	}
	if isbn != "" {
		query = query.Where("isbn = ?", isbn)
	}
	if author != "" {
		query = query.Where("author = ?", author)
	}
	if releaseDate != "" {
		query = query.Where("release_date = ?", releaseDate)
	}
	if availability != "" {
		query = query.Where("availability = ?", availability)
	}
	result := query.Find(&book)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Query failed")
	}
	return c.JSON(http.StatusOK, book)
}

func (h *Handler) getUsers(c echo.Context) error {
	var users []User

	tx := h.DB.Session(&gorm.Session{})
	result := tx.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Query failed")
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) getUser(c echo.Context) error {
	var user User

	id := c.QueryParam("id")
	log.Print(id)

	tx := h.DB

	if id == "" {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	res := tx.First(&user, id)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Query failed loser")
	}
	return c.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	os.Remove("./test.sqlite")
	db, err := gorm.Open(sqlite.Open("./test.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	h := Handler{DB: db}

	migrate(db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/Books", h.getBooks)
	e.GET("/Book", h.getBook)

	e.GET("/Users", h.getUsers)
	e.GET("/User", h.getUser)

	e.Logger.Fatal(e.Start(":1323"))
}
