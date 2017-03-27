package cause

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/auth"
	"politivate.org/web/models"
	"politivate.org/web/views"
)

func init() {
	mux[""] = whmux.Method{
		"GET":  http.HandlerFunc(cause),
		"POST": auth.WebLoginRequired(http.HandlerFunc(causeAction)),
	}
}

func cause(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	c := models.GetCause(ctx, causeId.MustGet(ctx))
	isAdministrating := false
	isFollowing := false
	if u != nil {
		isAdministrating = u.IsAdministrating(ctx, c)
		isFollowing = u.IsFollowing(ctx, c)
	}
	var challenges []*models.Challenge
	if isAdministrating {
		challenges = c.GetAllChallenges(ctx)
	} else {
		challenges = c.GetLiveChallenges(ctx)
	}
	views.Render(w, r, "cause", map[string]interface{}{
		"IsAdministrating": isAdministrating,
		"IsFollowing":      isFollowing,
		"Cause":            c,
		"Challenges":       challenges,
	})
}

func causeAction(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	switch r.FormValue("action") {
	case "delete":
		administerCause(r).Delete(ctx)
		whfatal.Redirect("/causes/")
	case "follow":
		auth.User(r).Follow(ctx, models.GetCause(ctx, causeId.MustGet(ctx)))
		whfatal.Redirect(r.RequestURI)
	case "unfollow":
		auth.User(r).Unfollow(ctx, models.GetCause(ctx, causeId.MustGet(ctx)))
		whfatal.Redirect(r.RequestURI)
	default:
		whfatal.Error(wherr.BadRequest.New("action not understood"))
	}
}
