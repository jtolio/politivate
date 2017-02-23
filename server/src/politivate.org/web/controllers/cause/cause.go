package cause

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/auth"
	"politivate.org/web/models"
	"politivate.org/web/views"
)

func init() {
	mux[""] = whmux.Method{
		"GET":  http.HandlerFunc(cause),
		"POST": http.HandlerFunc(editCause),
	}
}

func cause(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	c := models.GetCause(ctx, causeId.MustGet(ctx))
	isAdministrating := false
	if u != nil {
		isAdministrating = u.IsAdministrating(ctx, c)
	}
	views.Render(w, r, "cause", map[string]interface{}{
		"IsAdministrating": isAdministrating,
		"Cause":            c,
		"Challenges":       c.GetChallenges(ctx),
	})
}

func editCause(w http.ResponseWriter, r *http.Request) {
	switch r.FormValue("action") {
	case "delete":
		administerCause(r).Delete(whcompat.Context(r))
		whfatal.Redirect("/causes/")
	default:
		whfatal.Error(wherr.BadRequest.New("action not understood"))
	}
}
