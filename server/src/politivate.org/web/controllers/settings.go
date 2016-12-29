package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["settings"] = auth.WebLoginRequired(whmux.Exact(
		http.HandlerFunc(settings)))
}

func settings(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	Render(w, r, "settings", map[string]interface{}{
		"LogoutURL": auth.LogoutURL("/"),
		"User":      auth.User(ctx),
	})
}
