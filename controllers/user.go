package controllers

import (
	"azubi_library/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *Handler) GetUsers(c echo.Context) error {
	var users []models.User

	tx := h.DB.Session(&gorm.Session{})
	result := tx.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Query failed")
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUser(c echo.Context) error {
	var user models.User

	id := c.QueryParam("id")

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

func (h *Handler) CreateUser(c echo.Context) error {
	user := new(models.User)
	tx := h.DB.Session(&gorm.Session{})

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	tx.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
