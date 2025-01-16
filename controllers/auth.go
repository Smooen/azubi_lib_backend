package controllers

import "gorm.io/gorm"

type (
	Handler struct {
		DB *gorm.DB
	}
)

// func (h *Handler) Login() error {
//
// }
//
// func (h *Handler) Register() error {
//
// }
//
// func Logout() {
//
// }
