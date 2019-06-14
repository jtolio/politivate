package cause

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"
)

func init() {
	mux["challenges"] = whmux.Exact(
		http.HandlerFunc(serveCauseChallenges))
}

func serveCauseChallenges(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	whjson.Render(w, r, mustGetCause(ctx).GetLiveChallenges(ctx))
}
