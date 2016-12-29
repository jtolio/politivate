package cause

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"
)

var (
	challengeId = whmux.NewIntArg()
)

func init() {
	mux["challenge"] = challengeId.Shift(whmux.Exact(
		http.HandlerFunc(serveChallenge)))
}

func serveChallenge(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	whjson.Render(w, r,
		mustGetCause(ctx).GetChallenge(ctx, challengeId.MustGet(ctx)))
}
