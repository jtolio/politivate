package api

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/models"
)

func init() {
	mux["causes"] = whmux.Exact(http.HandlerFunc(serveCauses))
}

func serveCauses(w http.ResponseWriter, r *http.Request) {
	whjson.Render(w, r, models.GetCauses(whcompat.Context(r)))
}
