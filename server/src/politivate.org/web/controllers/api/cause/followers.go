package cause

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["followers"] = whmux.ExactPath(whmux.Method{
		"POST":   http.HandlerFunc(follow),
		"GET":    http.HandlerFunc(followers),
		"DELETE": http.HandlerFunc(unfollow),
	})
}

func follow(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	auth.User(ctx).Follow(ctx, mustGetCause(ctx))
	followers(w, r)
}

func unfollow(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	auth.User(ctx).Unfollow(ctx, mustGetCause(ctx))
	followers(w, r)
}

func followers(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(ctx)
	c := mustGetCause(ctx)
	whjson.Render(w, r, map[string]interface{}{
		"followers": c.UserCount(ctx),
		"following": u.IsFollowing(ctx, c),
	})
}
