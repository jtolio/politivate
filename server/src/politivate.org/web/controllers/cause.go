package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/models"
)

var (
	causeId = whmux.NewIntArg()
)

func init() {
	mux["cause"] = causeId.Shift(whmux.Exact(http.HandlerFunc(cause)))
}

func cause(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	Render(w, r, "cause", map[string]interface{}{
		"Cause": models.GetCause(ctx, causeId.MustGet(ctx))})
}
