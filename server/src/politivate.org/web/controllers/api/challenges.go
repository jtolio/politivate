package api

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/auth"
	"politivate.org/web/models"
)

func init() {
	authedMux["challenges"] = whmux.Exact(http.HandlerFunc(serveChallenges))
}

func serveChallenges(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	whjson.Render(w, r,
		models.GetChallenges(ctx, auth.User(r).CauseIds(ctx)...))
}
