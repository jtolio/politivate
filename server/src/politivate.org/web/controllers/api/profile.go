package api

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/auth"
)

func init() {
	authedMux["profile"] = whmux.Exact(http.HandlerFunc(serveProfile))
}

func serveProfile(w http.ResponseWriter, r *http.Request) {
	whjson.Render(w, r, auth.User(whcompat.Context(r)))
}
