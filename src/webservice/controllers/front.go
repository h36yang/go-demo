package controllers

import (
	"net/http"
)

// RegisterControllers initializes all controllers and HTTP routing
func RegisterControllers() {
	uc := newUserController()
	// Create routing to both /users and /users/{id}
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
