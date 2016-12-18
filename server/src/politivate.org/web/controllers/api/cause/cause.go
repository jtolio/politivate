package cause

import (
	"net/http"

	"github.com/jtolds/webhelp"
)

func init() {
	mux[""] = webhelp.RequireGet(http.HandlerFunc(serveCause))
}

func serveCause(w http.ResponseWriter, r *http.Request) {
	webhelp.RenderJSON(w, r, mustGetCause(webhelp.Context(r)))
}
