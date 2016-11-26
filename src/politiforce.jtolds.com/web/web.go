package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jtolds/webhelp"
	"golang.org/x/net/context"
)

type Cause struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Challenge struct {
	Cause      Cause  `json:"cause"`
	Title      string `json:"title"`
	ShortDesc  string `json:"short_desc"`
	PostedTS   int64  `json:"posted_ts"`
	DeadlineTS *int64 `json:"deadline_ts,omitempty"`
	Icon       string `json:"icon"`
	Points     int    `json:"points"`
}

func maybeInt64(val int64) *int64 { return &val }

var (
	now = time.Now()

	SierraClub = Cause{
		Name: "Sierra Club",
		Icon: "http://66.media.tumblr.com/avatar_cdbb9208e450_128.png"}

	challengeData = []Challenge{
		{
			Cause:      SierraClub,
			Title:      "Call your local representative",
			ShortDesc:  "We need you to tell them how important the environment is!",
			PostedTS:   now.UnixNano(),
			DeadlineTS: nil,
			Icon:       "http://www.iconsdb.com/icons/preview/black/office-phone-xxl.png",
			Points:     10,
		},
		{
			Cause:      SierraClub,
			Title:      "Show up to town hall",
			ShortDesc:  "We need you to tell them how important the environment is!",
			PostedTS:   now.UnixNano(),
			DeadlineTS: maybeInt64(now.UnixNano() + (7 * 24 * 60 * 60 * 1000000000)),
			Icon:       "https://cdn2.iconfinder.com/data/icons/the-urban-hustle-and-bustle/60/townhall-256.png",
			Points:     100,
		},
	}
)

func challenges(ctx context.Context,
	w webhelp.ResponseWriter, r *http.Request) error {
	data, err := json.MarshalIndent(map[string]interface{}{
		"response": challengeData,
	}, "", "  ")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	return err
}

func init() {
	http.Handle("/", webhelp.Base{Root: webhelp.DirMux{
		"challenges": webhelp.HandlerFunc(challenges),
	}})
}
