package controllers

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type (
	Handler struct {
		DB *gorm.DB
	}
)

type Claims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

// like this innit https://echo.labstack.com/docs/cookbook/jwt

// func (h *Handler) Login() error {
// 	// POST
// 	// gets username(== email), password form request body
// 	// if successfull then generate a jwt token and send it to cookie
// }
//
// func (h *Handler) Register() error {
// 	// gets email and hashed token from request body
// 	// if email is doesnt already exist
// 	// create new user
// }

// func Logout() error {
//
// }
