package api

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/models"
)

var (
	causeId = webhelp.NewIntArgMux()
)

func init() {
	mux["cause"] = causeId.Shift(webhelp.Exact(http.HandlerFunc(serveCause)))
	mux["causes"] = webhelp.Exact(http.HandlerFunc(serveCauses))
}

func serveCause(w http.ResponseWriter, r *http.Request) {
	c, err := models.GetCause(causeId.MustGet(webhelp.Context(r)))
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}
	webhelp.RenderJSON(w, r, c)
}

func serveCauses(w http.ResponseWriter, r *http.Request) {
	c, err := models.GetCauses()
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}
	webhelp.RenderJSON(w, r, c)
}
