package models

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/whfatal"
)

type Action struct {
	Id     int64 `json:"-", datastore:"-"`
	UserId int64 `json:"user_id" datastore:"-"`

	CauseId     int64 `json:"cause_id"`
	ChallengeId int64 `json:"challenge_id"`
	When        Time  `json:"when"`

	// set if the challenge was a location challenge
	// the latitude/longitude of the challenge completed, not of the user
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	// set if the challenge was a phonecall challenge
	// the phone number the user called
	Phone string `json:"phone"`
}

func (chal *Challenge) Action(ctx context.Context, u *User) *Action {
	return &Action{
		UserId:      u.Id,
		CauseId:     chal.CauseId,
		ChallengeId: chal.Id,
		When:        TimeNow(),
	}
}

func (a *Action) Save(ctx context.Context) {
	if a.UserId == 0 || a.CauseId == 0 || a.ChallengeId == 0 ||
		a.When.Time.IsZero() {
		whfatal.Error(Error.New("incomplete action"))
	}
	k, err := datastore.Put(ctx,
		datastore.NewKey(ctx, "Action", "", a.Id, userKey(ctx, a.UserId)), a)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	a.Id = k.IntID()
}

func (chal *Challenge) Completed(ctx context.Context, u *User,
	after time.Time) []*Action {
	// use make so the json doesn't look like `null`
	actions := make([]*Action, 0)

	if chal.Id == 0 || chal.CauseId == 0 || u.Id == 0 {
		return actions
	}

	keys, err := datastore.NewQuery("Action").Ancestor(userKey(ctx, u.Id)).
		Filter("CauseId =", chal.CauseId).Filter("ChallengeId =", chal.Id).
		Filter("When.Time >=", after).GetAll(ctx, &actions)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}

	for i, key := range keys {
		actions[i].Id = key.IntID()
		actions[i].UserId = u.Id
	}
	return actions
}
