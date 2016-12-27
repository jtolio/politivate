package api

import (
	"net/http"

	"github.com/jtolds/webhelp/whcompat"
	"github.com/jtolds/webhelp/whjson"
	"github.com/jtolds/webhelp/whmux"

	"politivate.org/web/models"
)

func init() {
	mux["causes"] = whmux.Exact(http.HandlerFunc(serveCauses))
}

func serveCauses(w http.ResponseWriter, r *http.Request) {
	whjson.Render(w, r, models.GetCauses(whcompat.Context(r)))
}
