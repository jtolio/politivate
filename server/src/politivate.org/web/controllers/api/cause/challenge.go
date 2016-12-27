package cause

import (
	"net/http"

	"github.com/jtolds/webhelp/whcompat"
	"github.com/jtolds/webhelp/whjson"
	"github.com/jtolds/webhelp/whmux"
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
