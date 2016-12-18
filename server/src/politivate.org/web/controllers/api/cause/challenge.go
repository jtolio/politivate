package cause

import (
	"net/http"

	"github.com/jtolds/webhelp"
)

var (
	challengeId = webhelp.NewIntArgMux()
)

func init() {
	mux["challenge"] = challengeId.Shift(webhelp.Exact(
		http.HandlerFunc(serveChallenge)))
}

func serveChallenge(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	webhelp.RenderJSON(w, r,
		mustGetCause(ctx).GetChallenge(ctx, challengeId.MustGet(ctx)))
}
