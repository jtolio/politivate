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

func init() {
	mux["causes"] = whmux.Dir{
		"": whmux.Exact(http.HandlerFunc(causes)),
		"new": whmux.Method{
			"GET":  http.HandlerFunc(newCauseForm),
			"POST": http.HandlerFunc(newCauseCreation),
		},
	}
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
		"Error": "",
		"Form":  map[string]string{},
	})
}

func newCauseCreation(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	if u == nil || !u.CanCreateCause {
		whfatal.Redirect("/causes/")
	}

	c := models.NewCause(ctx)
	c.Name = r.FormValue("name")
	c.URL = r.FormValue("url")
	c.Description = r.FormValue("description")
	if c.Name == "" || c.URL == "" || c.Description == "" {
		Render(w, r, "new_cause", map[string]interface{}{
			"Error": "Required field missing",
			"Form": map[string]string{
				"name":        c.Name,
				"url":         c.URL,
				"description": c.Description,
			},
		})
		return
	}

	c.Save(ctx)
	u.Administrate(ctx, c)

	whfatal.Redirect(fmt.Sprintf("/cause/%d", c.Id))
}
