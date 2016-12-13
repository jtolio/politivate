package models

import (
	"time"
)

type Challenge struct {
	Id         int64  `json:"id"`
	CauseId    int64  `json:"cause_id"`
	Title      string `json:"title"`
	ShortDesc  string `json:"short_desc"`
	PostedTS   int64  `json:"posted_ts"`
	DeadlineTS *int64 `json:"deadline_ts,omitempty"`
	IconURL    string `json:"icon_url"`
	Points     int    `json:"points"`
}

func GetChallenge(id int64) (*Challenge, error) {
	for _, c := range TESTCHALLENGES {
		if c.Id == id {
			return c, nil
		}
	}
	return nil, NotFound.New("challenge %d not found", id)
}

func GetChallenges() ([]*Challenge, error) {
	return TESTCHALLENGES, nil
}

// TEST DATA

func maybeInt64(val int64) *int64 { return &val }

var TESTCHALLENGES = []*Challenge{
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
