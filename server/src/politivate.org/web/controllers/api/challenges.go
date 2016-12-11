package api

import (
	"net/http"
	"time"

	"github.com/jtolds/webhelp"

	"politivate.org/web/models"
)

var (
	challengeId = webhelp.NewIntArgMux()
)

func init() {
	Mux["challenge"] = challengeId.Shift(webhelp.Exact(
		http.HandlerFunc(serveChallenge)))
	Mux["challenges"] = webhelp.Exact(http.HandlerFunc(serveChallenges))
}

func serveChallenge(w http.ResponseWriter, r *http.Request) {
	id := challengeId.MustGet(webhelp.Context(r))
	for _, c := range TESTCHALLENGES {
		if c.Id == id {
			RenderJSON(w, r, c)
			return
		}
	}
	webhelp.HandleError(w, r,
		webhelp.ErrNotFound.New("challenge %d not found", id))
}

func serveChallenges(w http.ResponseWriter, r *http.Request) {
	RenderJSON(w, r, TESTCHALLENGES)
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
