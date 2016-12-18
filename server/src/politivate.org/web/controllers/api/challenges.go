package api

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/models"
)

func init() {
	mux["challenges"] = webhelp.Exact(http.HandlerFunc(serveChallenges))
}

func serveChallenges(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	challenges := make([]*models.Challenge, 0) // so the json isn't `null`
	for _, cause := range models.GetCauses(ctx) {
		challenges = append(challenges, cause.GetChallenges(ctx)...)
	}
	webhelp.RenderJSON(w, r, challenges)
}
