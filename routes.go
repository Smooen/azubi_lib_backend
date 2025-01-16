package main

import (
	"azubi_library/controllers"
	"github.com/labstack/echo/v4"
)



func SetupRoutes(e *echo.Echo, h *controllers.Handler) {
	// probably combine into /books -> just take query params if there are any, if not return all books
	e.GET("/books", h.GetBooks)
	e.GET("/book", h.GetBook)
	// e.GET("/Login", h.Login)
	// e.GET("/Register", h.Register)
}
