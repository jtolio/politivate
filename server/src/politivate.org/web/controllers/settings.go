package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp/whcompat"
	"github.com/jtolds/webhelp/whmux"

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
