package api

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/auth"
)

func init() {
	authedMux["challenges"] = whmux.Exact(http.HandlerFunc(serveChallenges))
}

func serveChallenges(w http.ResponseWriter, r *http.Request) {
	whjson.Render(w, r, auth.User(r).GetLiveChallenges(whcompat.Context(r)))
}
