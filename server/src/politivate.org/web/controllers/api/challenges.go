package api

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/controllers/auth"
	"politivate.org/web/models"
)

func init() {
	mux["challenges"] = webhelp.Exact(http.HandlerFunc(serveChallenges))
}

func serveChallenges(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	webhelp.RenderJSON(w, r,
		models.GetChallenges(ctx, auth.User(ctx).CauseIds(ctx)...))
}
