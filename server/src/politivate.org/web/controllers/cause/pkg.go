package cause

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
	causeId     = whmux.NewIntArg()
	challengeId = whmux.NewIntArg()
)

var (
	mux = whmux.Dir{}

	Handler http.Handler = causeId.Shift(mux)
)

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

func administerChallenge(r *http.Request) *models.Challenge {
	ctx := whcompat.Context(r)
	return administerCause(r).GetChallenge(ctx, challengeId.MustGet(ctx))
}
