package controllers

import (
	"net/http"
)

func init() {
	Mux["settings"] = LoginRequired(http.HandlerFunc(settings))
}

func settings(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "settings", map[string]interface{}{
		"LogoutURL": Auth.LogoutAllURL("/")})
}
