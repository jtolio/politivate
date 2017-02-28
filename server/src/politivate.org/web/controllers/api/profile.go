package api

import (
	"net/http"
	"time"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/auth"
)

func init() {
	authedMux["profile"] = whmux.Exact(http.HandlerFunc(serveProfile))
}

func serveProfile(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	m := u.JSON()
	m["month_actions"] = u.Actions(ctx, time.Now().Add(-30*24*time.Hour))
	whjson.Render(w, r, m)
}
