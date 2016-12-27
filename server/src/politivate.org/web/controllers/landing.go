package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp/whmux"
)

func init() {
	mux[""] = whmux.RequireGet(http.HandlerFunc(landing))
}

func landing(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "landing", map[string]interface{}{})
}
