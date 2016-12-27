package cause

import (
	"net/http"

	"github.com/jtolds/webhelp/whcompat"
	"github.com/jtolds/webhelp/whjson"
	"github.com/jtolds/webhelp/whmux"
)

func init() {
	mux["challenges"] = whmux.Exact(
		http.HandlerFunc(serveCauseChallenges))
}

func serveCauseChallenges(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	whjson.Render(w, r, mustGetCause(ctx).GetChallenges(ctx))
}
