package api

import (
	"net/http"

	"github.com/jtolds/webhelp/whcompat"
	"github.com/jtolds/webhelp/whjson"
	"github.com/jtolds/webhelp/whmux"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["profile"] = whmux.Exact(http.HandlerFunc(serveProfile))
}

func serveProfile(w http.ResponseWriter, r *http.Request) {
	whjson.Render(w, r, auth.User(whcompat.Context(r)))
}
