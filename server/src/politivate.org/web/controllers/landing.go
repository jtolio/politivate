package controllers

import (
	"net/http"
)

func init() {
	Mux[""] = http.HandlerFunc(landing)
}

func landing(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "landing", map[string]interface{}{})
}
