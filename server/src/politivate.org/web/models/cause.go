package models

type Cause struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
}

func GetCause(id int64) (*Cause, error) {
	for _, c := range TESTCAUSES {
		if c.Id == id {
			return c, nil
		}
	}
	return nil, NotFound.New("cause %d not found", id)
}

func GetCauses() ([]*Cause, error) {
	return TESTCAUSES, nil
}

// TEST DATA

var TESTCAUSES = []*Cause{
	{
		Id:      1,
		Name:    "Sierra Club",
		IconURL: "http://66.media.tumblr.com/avatar_cdbb9208e450_128.png",
	},
}
