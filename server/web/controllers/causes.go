package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/auth"
	"politivate.org/web/forms"
	"politivate.org/web/models"
)

func init() {
	mux["causes"] = Beta(whmux.Dir{
		"": whmux.Exact(http.HandlerFunc(causes)),
		"new": whmux.ExactPath(whmux.Method{
			"GET":  http.HandlerFunc(newCauseForm),
			"POST": http.HandlerFunc(newCauseCreation),
		}),
	})
}

func causes(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "causes", map[string]interface{}{
		"Causes": models.GetCauses(whcompat.Context(r)),
	})
}

func newCauseForm(w http.ResponseWriter, r *http.Request) {
	u := auth.User(r)
	if u == nil || !u.CanCreateCause {
		whfatal.Redirect("/causes/")
	}

	Render(w, r, "new_cause", map[string]interface{}{
		"Form": forms.NewCauseForm(),
	})
}

func newCauseCreation(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	if u == nil || !u.CanCreateCause {
		whfatal.Redirect("/causes/")
	}

	c := models.NewCause(ctx)
	ok, f := forms.ProcessCauseForm(c, r)
	if !ok {
		Render(w, r, "new_cause", map[string]interface{}{
			"Form": f,
		})
		return
	}
	c.Save(ctx)
	u.Administrate(ctx, c)

	whfatal.Redirect(fmt.Sprintf("/cause/%d", c.Id))
}
