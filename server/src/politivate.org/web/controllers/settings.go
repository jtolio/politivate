package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["settings"] = auth.WebLoginRequired(webhelp.Exact(
		http.HandlerFunc(settings)))
}

func settings(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "settings", map[string]interface{}{
		"LogoutURL": auth.LogoutURL("/")})
}
