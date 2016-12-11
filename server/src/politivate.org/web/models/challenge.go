package models

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
