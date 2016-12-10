package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jtolds/webhelp"
)

func init() {
	Mux["challenges"] = http.HandlerFunc(challenges)
}

type Cause struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Challenge struct {
	Id         int    `json:"id"`
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
		Id:   1,
		Name: "Sierra Club",
		Icon: "http://66.media.tumblr.com/avatar_cdbb9208e450_128.png"}

	challengeData = []Challenge{
		{
			Id:         2,
			Cause:      SierraClub,
			Title:      "Call your local representative",
			ShortDesc:  "We need you to tell them how important the environment is!",
			PostedTS:   now.UnixNano(),
			DeadlineTS: nil,
			Icon:       "http://www.iconsdb.com/icons/preview/black/office-phone-xxl.png",
			Points:     10,
		},
		{
			Id:         3,
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

func challenges(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(map[string]interface{}{
		"response": challengeData,
	}, "", "  ")
	if err != nil {
		webhelp.FatalError(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	webhelp.FatalError(err)
}
