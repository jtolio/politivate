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
	mux["challenge"] = challengeId.Shift(webhelp.Exact(
		http.HandlerFunc(serveChallenge)))
	mux["challenges"] = webhelp.Exact(http.HandlerFunc(serveChallenges))
}

func serveChallenge(w http.ResponseWriter, r *http.Request) {
	id := challengeId.MustGet(webhelp.Context(r))
	for _, c := range TESTCHALLENGES {
		if c.Id == id {
			webhelp.RenderJSON(w, r, c)
			return
		}
	}
	webhelp.HandleError(w, r,
		webhelp.ErrNotFound.New("challenge %d not found", id))
}

func serveChallenges(w http.ResponseWriter, r *http.Request) {
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

	webhelp.RenderJSON(w, r, append([]models.Challenge{
		{
			Id:         4,
			CauseId:    1,
			Title:      info.Name,
			ShortDesc:  info.Link,
			PostedTS:   time.Now().UnixNano(),
			DeadlineTS: nil,
			IconURL:    info.Picture,
			Points:     10,
		},
	}, TESTCHALLENGES...))
}

// TEST DATA

func maybeInt64(val int64) *int64 { return &val }

var TESTCHALLENGES = []models.Challenge{
	{
		Id:         2,
		CauseId:    1,
		Title:      "Call your local representative",
		ShortDesc:  "We need you to tell them how important the environment is!",
		PostedTS:   time.Now().UnixNano(),
		DeadlineTS: nil,
		IconURL:    "http://www.iconsdb.com/icons/preview/black/office-phone-xxl.png",
		Points:     10,
	},
	{
		Id:         3,
		CauseId:    1,
		Title:      "Show up to town hall",
		ShortDesc:  "We need you to tell them how important the environment is!",
		PostedTS:   time.Now().UnixNano(),
		DeadlineTS: maybeInt64(time.Now().UnixNano() + (7 * 24 * 60 * 60 * 1000000000)),
		IconURL:    "https://cdn2.iconfinder.com/data/icons/the-urban-hustle-and-bustle/60/townhall-256.png",
		Points:     100,
	},
}
