package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jtolds/webhelp"
	"golang.org/x/oauth2"

	"politivate.org/web/controllers/auth"
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

	// stupidness -------------------------------------
	// TODO: this is just a bunch of stupidness to try and make sure that auth
	// works. returning a challenge with the user's information makes no sense,
	// but that way i don't have to change the app much to test it.
	if r.FormValue("testauth") != "" {
		provider, ok := auth.Auth.Handler("google")
		if !ok {
			webhelp.FatalError(webhelp.ErrInternalServerError.New("uh oh"))
		}
		c := provider.Provider().Client(webhelp.Context(r),
			&oauth2.Token{AccessToken: r.Header.Get("X-Auth-Token-Google")})

		resp, err := c.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
		if err != nil {
			webhelp.FatalError(err)
		}
		defer resp.Body.Close()

		var info struct {
			Id         string `json:"id"`
			Name       string `json:"name"`
			GivenName  string `json:"given_name"`
			FamilyName string `json:"family_name"`
			Link       string `json:"link"`
			Picture    string `json:"picture"`
			Gender     string `json:"gender"`
			Locale     string `json:"locale"`
		}

		err = json.NewDecoder(resp.Body).Decode(&info)
		if err != nil {
			webhelp.FatalError(err)
		}

		challenges = append(challenges, &models.Challenge{
			Id:        -1,
			CauseId:   -1,
			Title:     info.Name,
			ShortDesc: info.Link,
			Posted:    time.Now(),
			IconURL:   info.Picture,
			Points:    10,
		})
	}
	// okay, end stupidness ----------------------

	webhelp.RenderJSON(w, r, challenges)
}
