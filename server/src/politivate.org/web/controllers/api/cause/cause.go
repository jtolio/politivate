package cause

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"
)

func init() {
	mux[""] = whmux.RequireGet(http.HandlerFunc(serveCause))
}

func serveCause(w http.ResponseWriter, r *http.Request) {
	whjson.Render(w, r, mustGetCause(whcompat.Context(r)))
}
