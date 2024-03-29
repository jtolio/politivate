package cause

import (
	"fmt"
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/auth"
	"politivate.org/web/models"
	"politivate.org/web/views"
)

var (
	challengeMux = whmux.Dir{
		"": whmux.Method{
			"GET":  http.HandlerFunc(challenge),
			"POST": auth.WebLoginRequired(http.HandlerFunc(challengeAction)),
		}}
)

func init() {
	mux["challenge"] = challengeId.Shift(challengeMux)
}

func challenge(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	cause := models.GetCause(ctx, causeId.MustGet(ctx))
	isAdministrating := false
	if u != nil {
		isAdministrating = u.IsAdministrating(ctx, cause)
	}
	views.Render(w, r, "challenge", map[string]interface{}{
		"IsAdministrating": isAdministrating,
		"Cause":            cause,
		"Challenge":        cause.GetChallenge(ctx, challengeId.MustGet(ctx)),
	})
}

func challengeAction(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	switch r.FormValue("action") {
	case "delete":
		administerChallenge(r).Delete(ctx)
		whfatal.Redirect(fmt.Sprintf("/cause/%d", causeId.MustGet(ctx)))
	case "enable":
		chal := administerChallenge(r)
		chal.Info.Enabled = true
		chal.Save(ctx)
		whfatal.Redirect(r.RequestURI)
	case "disable":
		chal := administerChallenge(r)
		chal.Info.Enabled = false
		chal.Save(ctx)
		whfatal.Redirect(r.RequestURI)
	default:
		whfatal.Error(wherr.BadRequest.New("action not understood"))
	}
}
