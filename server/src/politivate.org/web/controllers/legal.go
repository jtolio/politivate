package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whmux"
)

func init() {
	mux["legal"] = whmux.Dir{
		"tos":     whmux.Exact(http.HandlerFunc(tos)),
		"privacy": whmux.Exact(http.HandlerFunc(privacy)),
	}
}

func tos(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "tos", map[string]interface{}{})
}

func privacy(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "privacy", map[string]interface{}{})
}
