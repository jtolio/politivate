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
				"GET":  http.HandlerFunc(newChallengeForm),
				"POST": http.HandlerFunc(newChallengeCreate),
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

func administerCause(r *http.Request) *models.Cause {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	c := models.GetCause(ctx, causeId.MustGet(ctx))
	if u == nil || !u.IsAdministrating(ctx, c) {
		whfatal.Redirect(fmt.Sprintf("/cause/%d", c.Id))
		return nil
	}
	return c
}

func newChallengeForm(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "new_challenge", map[string]interface{}{
		"Cause": administerCause(r),
		"Form":  map[string]interface{}{},
	})
}

func newChallengeCreate(w http.ResponseWriter, r *http.Request) {
	c := administerCause(r)
	Render(w, r, "new_challenge", map[string]interface{}{
		"Cause": c,
		"Form":  map[string]interface{}{},
	})
}
