package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers initializes all controllers and HTTP routing
func RegisterControllers() {
	uc := newUserController()
	// Create routing to both /users and /users/{id}
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
