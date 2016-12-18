package api

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["profile"] = webhelp.Exact(http.HandlerFunc(serveProfile))
}

func serveProfile(w http.ResponseWriter, r *http.Request) {
	webhelp.RenderJSON(w, r, auth.User(webhelp.Context(r)))
}
