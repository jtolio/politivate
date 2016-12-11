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
	Mux["cause"] = causeId.Shift(webhelp.Exact(http.HandlerFunc(serveCause)))
	Mux["causes"] = webhelp.Exact(http.HandlerFunc(serveCauses))
}

func serveCause(w http.ResponseWriter, r *http.Request) {
	id := causeId.MustGet(webhelp.Context(r))
	for _, c := range TESTCAUSES {
		if c.Id == id {
			webhelp.RenderJSON(w, r, c)
			return
		}
	}
	webhelp.HandleError(w, r, webhelp.ErrNotFound.New("cause %d not found", id))
}

func serveCauses(w http.ResponseWriter, r *http.Request) {
	webhelp.RenderJSON(w, r, TESTCAUSES)
}

// TEST DATA

var TESTCAUSES = []models.Cause{{
	Id:      1,
	Name:    "Sierra Club",
	IconURL: "http://66.media.tumblr.com/avatar_cdbb9208e450_128.png"}}
