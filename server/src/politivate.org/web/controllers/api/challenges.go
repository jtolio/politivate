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
	mux["challenges"] = webhelp.Exact(http.HandlerFunc(serveChallenges))
}

func serveChallenge(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	cause, err := getCause(ctx)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}
	challenge, err := cause.GetChallenge(ctx, challengeId.MustGet(ctx))
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}
	webhelp.RenderJSON(w, r, challenge)
}

func serveChallenges(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	challenges := make([]*models.Challenge, 0) // so the json isn't `null`
	causes, err := models.GetCauses(ctx)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}
	for _, cause := range causes {
		causeChallenges, err := cause.GetChallenges(ctx)
		if err != nil {
			webhelp.HandleError(w, r, err)
			return
		}
		challenges = append(challenges, causeChallenges...)
	}

	// stupidness -------------------------------------
	// TODO: this is just a bunch of stupidness to try and make sure that auth
	// works. returning a challenge with the user's information makes no sense,
	// but that way i don't have to change the app much to test it.
	if r.FormValue("testauth") != "" {
		provider, ok := auth.Auth.Handler("google")
		if !ok {
			webhelp.HandleError(w, r, webhelp.ErrInternalServerError.New("uh oh"))
			return
		}
		c := provider.Provider().Client(webhelp.Context(r),
			&oauth2.Token{AccessToken: r.Header.Get("X-Auth-Token-Google")})

		resp, err := c.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
		if err != nil {
			webhelp.HandleError(w, r, err)
			return
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
			webhelp.HandleError(w, r, err)
			return
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
