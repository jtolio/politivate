package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whmux"
)

func init() {
	mux["about"] = whmux.RequireGet(http.HandlerFunc(about))
}

func about(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "about", map[string]interface{}{})
}
