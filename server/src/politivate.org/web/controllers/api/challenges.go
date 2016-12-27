package api

import (
	"net/http"

	"github.com/jtolds/webhelp/whcompat"
	"github.com/jtolds/webhelp/whjson"
	"github.com/jtolds/webhelp/whmux"

	"politivate.org/web/controllers/auth"
	"politivate.org/web/models"
)

func init() {
	mux["challenges"] = whmux.Exact(http.HandlerFunc(serveChallenges))
}

func serveChallenges(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	whjson.Render(w, r,
		models.GetChallenges(ctx, auth.User(ctx).CauseIds(ctx)...))
}
