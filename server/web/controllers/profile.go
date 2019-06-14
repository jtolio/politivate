package controllers

import (
	"net/http"
	"time"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/auth"
)

func init() {
	mux["profile"] = auth.WebLoginRequired(Beta(whmux.Exact(http.HandlerFunc(
		profileHandler))))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	Render(w, r, "profile", map[string]interface{}{
		"User":    u,
		"Actions": u.Actions(ctx, time.Time{}),
	})
}
