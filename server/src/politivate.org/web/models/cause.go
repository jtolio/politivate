package models

type Cause struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
}
