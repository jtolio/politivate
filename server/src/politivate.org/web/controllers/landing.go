package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whmux"
)

func init() {
	mux[""] = whmux.RequireGet(http.HandlerFunc(landing))
}

func landing(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "landing", map[string]interface{}{})
}
