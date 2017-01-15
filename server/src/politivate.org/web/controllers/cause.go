package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/auth"
	"politivate.org/web/models"
)

var (
	causeId = whmux.NewIntArg()
)

func init() {
	mux["cause"] = causeId.Shift(whmux.Dir{
		"": whmux.Exact(http.HandlerFunc(cause)),
		"challenges": whmux.Dir{
			"new": whmux.Method{
				"GET": http.HandlerFunc(newChallenge),
			},
		},
	})
}

func cause(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	c := models.GetCause(ctx, causeId.MustGet(ctx))
	isAdministrating := false
	if u != nil {
		isAdministrating = u.IsAdministrating(ctx, c)
	}
	Render(w, r, "cause", map[string]interface{}{
		"IsAdministrating": isAdministrating,
		"Cause":            c,
	})
}

func newChallenge(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	c := models.GetCause(ctx, causeId.MustGet(ctx))
	if u == nil || !u.IsAdministrating(ctx, c) {
		whfatal.Redirect(fmt.Sprintf("/cause/%d", c.Id))
		return
	}

	Render(w, r, "new_challenge", map[string]interface{}{
		"Cause": c,
		"Form":  map[string]interface{}{},
	})
}
