package api

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/models"
)

func init() {
	mux["causes"] = webhelp.Exact(http.HandlerFunc(serveCauses))
}

func serveCauses(w http.ResponseWriter, r *http.Request) {
	webhelp.RenderJSON(w, r, models.GetCauses(webhelp.Context(r)))
}
