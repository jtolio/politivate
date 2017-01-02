package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whmux"
)

func init() {
	mux["get"] = whmux.Exact(http.HandlerFunc(get))
}

func get(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "get", map[string]interface{}{})
}
