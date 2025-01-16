package controllers

import (
	"azubi_library/models"
	"log"
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
	log.Print(id)

	tx := h.DB

	if id == "" {
		return c.JSON(http.StatusNotFound, "Book not found")
	}

	res := tx.First(&book, id)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Query failed loser")
	}

	return c.JSON(http.StatusOK, book)
}
