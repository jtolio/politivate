package cause

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["follow"] = webhelp.ExactPath(webhelp.RequireMethod("POST",
		http.HandlerFunc(follow)))
	mux["unfollow"] = webhelp.ExactPath(webhelp.RequireMethod("POST",
		http.HandlerFunc(unfollow)))
}

func follow(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	auth.User(ctx).Follow(ctx, mustGetCause(ctx))
	webhelp.RenderJSON(w, r, "ok")
}

func unfollow(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	auth.User(ctx).Unfollow(ctx, mustGetCause(ctx))
	webhelp.RenderJSON(w, r, "ok")
}
