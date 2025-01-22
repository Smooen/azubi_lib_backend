package controllers

import (
	"azubi_library/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	Handler struct {
		DB *gorm.DB
	}
)

type jwtCustomClaims struct {
	Username  string `json:"username"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func (h *Handler) Login(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, "Couldnt bind user to echo context")
	}

	var existingUser models.User

	res := h.DB.Find(&existingUser, "username = ?", user.Username)

	if res.Error != nil {
		c.JSON(http.StatusUnauthorized, "User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password));
	err != nil {
		return c.JSON(http.StatusUnauthorized, "Password incorrect")
	}

	// generate token and put into cookie jar

	return c.Redirect(http.StatusMovedPermanently, "/books")
}

//
// func (h *Handler) Register() error {
// 	// gets email and hashed token from request body
// 	// if email is doesnt already exist
// 	// create new user
// }

// func Logout() error {
//
// }
