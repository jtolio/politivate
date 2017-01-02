package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/models"
)

func init() {
	mux["causes"] = whmux.Exact(http.HandlerFunc(causes))
}

func causes(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "causes", map[string]interface{}{
		"Causes": models.GetCauses(whcompat.Context(r))})
}
