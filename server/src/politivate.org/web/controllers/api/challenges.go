package api

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/models"
)

var (
	challengeId = webhelp.NewIntArgMux()
)

func init() {
	causeMux["challenge"] = challengeId.Shift(webhelp.Exact(
		http.HandlerFunc(serveChallenge)))
	causeMux["challenges"] = webhelp.Exact(
		http.HandlerFunc(serveCauseChallenges))

	mux["challenges"] = webhelp.Exact(http.HandlerFunc(serveChallenges))
}

func serveChallenge(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	webhelp.RenderJSON(w, r,
		mustGetCause(ctx).GetChallenge(ctx, challengeId.MustGet(ctx)))
}

func serveCauseChallenges(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	webhelp.RenderJSON(w, r, mustGetCause(ctx).GetChallenges(ctx))
}

func serveChallenges(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	challenges := make([]*models.Challenge, 0) // so the json isn't `null`
	for _, cause := range models.GetCauses(ctx) {
		challenges = append(challenges, cause.GetChallenges(ctx)...)
	}
	webhelp.RenderJSON(w, r, challenges)
}
