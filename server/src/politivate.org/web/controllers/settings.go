package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp"
)

func init() {
	Mux["settings"] = LoginRequired(webhelp.Exact(http.HandlerFunc(settings)))
}

func settings(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "settings", map[string]interface{}{
		"LogoutURL": Auth.LogoutAllURL("/")})
}
