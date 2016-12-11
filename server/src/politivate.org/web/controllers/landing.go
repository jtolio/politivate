package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp"
)

func init() {
	Mux[""] = webhelp.RequireGet(http.HandlerFunc(landing))
}

func landing(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "landing", map[string]interface{}{})
}