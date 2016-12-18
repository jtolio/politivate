package cause

import (
	"net/http"

	"github.com/jtolds/webhelp"
)

func init() {
	mux["challenges"] = webhelp.Exact(
		http.HandlerFunc(serveCauseChallenges))
}

func serveCauseChallenges(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	webhelp.RenderJSON(w, r, mustGetCause(ctx).GetChallenges(ctx))
}
