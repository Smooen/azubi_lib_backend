package controllers

import (
	"azubi_library/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)


func (h *Handler) GetBooks(c echo.Context) error {
	var books []models.Book

	tx := h.DB.Session(&gorm.Session{})
	result := tx.Find(&books)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Query failed")
	}

	return c.JSON(http.StatusOK, books);
}

func (h *Handler) GetBook(c echo.Context) error {
	var book models.Book

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
