package api

import (
	"net/http"
	"time"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/auth"
)

func init() {
	authedMux["profile"] = whmux.Exact(http.HandlerFunc(serveProfile))
}

func serveProfile(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	m := u.JSON()
	actions := u.Actions(ctx, time.Time{})
	m["month_actions"] = actions
	m["actions"] = actions
	whjson.Render(w, r, m)
}
