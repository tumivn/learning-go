package handlers

import (
	"github.com/tumivn/go-booking-app/pkg/render"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "About", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "Home", nil)
}
