package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp"
)

func init() {
	Mux["legal"] = webhelp.DirMux{
		"tos":     webhelp.Exact(http.HandlerFunc(tos)),
		"privacy": webhelp.Exact(http.HandlerFunc(privacy)),
	}
}

func tos(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "tos", map[string]interface{}{})
}

func privacy(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "privacy", map[string]interface{}{})
}
