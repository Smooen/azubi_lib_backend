package controllers

import (
	"azubi_library/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	Handler struct {
		DB *gorm.DB
	}
)

type jwtCustomClaims struct {
	Username string `json:"username"`
	//Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func generateToken(user *models.User, c echo.Context) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Hour * 72)
	claims := &jwtCustomClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("giga-secret")) // secret should be key from env
	if err != nil {
		c.JSON(500, "Could not generate token")
		return "", expirationTime, err
	}

	return tokenString, expirationTime, nil
}

func setTokenCookie(c echo.Context, token string, expirationTime time.Time) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = expirationTime
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func (h *Handler) Login(c echo.Context) error {
	var loginReq LoginRequest

	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusInternalServerError, "Couldnt bind user to echo context")
	}

	var existingUser models.User

	res := h.DB.Find(&existingUser, "email = ?", loginReq.Email)

	if res.Error != nil {
		c.JSON(http.StatusUnauthorized, "User not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginReq.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	log.Info("Generating token for user: ", loginReq.Email)
	token, expirationTime, err := generateToken(&existingUser, c)
	if err != nil {
		return c.JSON(500, "Something went wrong creating the token")
	}
	setTokenCookie(c, token, expirationTime)

	return c.Redirect(http.StatusMovedPermanently, "/books")
}

func (h *Handler) Register(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, "Couldnt bind user to echo context")
	}

	var existingUser models.User
	if err := h.DB.Where("email = ? OR username = ?", user.Email, user.Username).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusBadRequest, "User already exists")
	}

	pwHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	user.Password = string(pwHash)

	if err := h.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Couldn't create user")
	}

	token, expirationTime, err := generateToken(&user, c)
	if err != nil {
		return c.JSON(500, "Something went wrong creating the token")
	}
	setTokenCookie(c, token, expirationTime)

	return c.Redirect(http.StatusCreated, "/books")
}

// func Logout() error {
//
// }
