package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/controllers/auth"
)

func init() {
	Mux["settings"] = auth.LoginRequired(webhelp.Exact(
		http.HandlerFunc(settings)))
}

func settings(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "settings", map[string]interface{}{
		"LogoutURL": auth.Auth.LogoutAllURL("/")})
}
