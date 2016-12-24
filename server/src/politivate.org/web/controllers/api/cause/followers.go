package cause

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["followers"] = webhelp.ExactPath(webhelp.MethodMux{
		"POST":   http.HandlerFunc(follow),
		"GET":    http.HandlerFunc(followers),
		"DELETE": http.HandlerFunc(unfollow),
	})
}

func follow(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	auth.User(ctx).Follow(ctx, mustGetCause(ctx))
	followers(w, r)
}

func unfollow(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	auth.User(ctx).Unfollow(ctx, mustGetCause(ctx))
	followers(w, r)
}

func followers(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	u := auth.User(ctx)
	c := mustGetCause(ctx)
	webhelp.RenderJSON(w, r, map[string]interface{}{
		"followers": c.UserCount(ctx),
		"following": u.IsFollowing(ctx, c),
	})
}
