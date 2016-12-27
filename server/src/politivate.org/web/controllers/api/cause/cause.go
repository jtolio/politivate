package cause

import (
	"net/http"

	"github.com/jtolds/webhelp/whcompat"
	"github.com/jtolds/webhelp/whjson"
	"github.com/jtolds/webhelp/whmux"
)

func init() {
	mux[""] = whmux.RequireGet(http.HandlerFunc(serveCause))
}

func serveCause(w http.ResponseWriter, r *http.Request) {
	whjson.Render(w, r, mustGetCause(whcompat.Context(r)))
}
