package models

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/whfatal"
)

const (
	// We want to tell users that they can only complete a challenge once every
	// day, so that's what we'll message, but we only limit to once every 18 hours.
	// This way, if someone calls once late in the day and wants to call earlier
	// the next day, they can.
	MinChallengeInterval = time.Hour * 18
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

func (a *Action) Cause(ctx context.Context) *Cause {
	return GetCause(ctx, a.CauseId)
}

func (a *Action) Challenge(ctx context.Context) *Challenge {
	return GetChallenge(ctx, a.ChallengeId, a.CauseId)
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

func (u *User) getActions(ctx context.Context,
	q func(*datastore.Query) *datastore.Query) []*Action {
	// use make so the json doesn't look like `null`
	actions := make([]*Action, 0)
	if u.Id == 0 {
		return actions
	}
	keys, err := q(datastore.NewQuery("Action").Ancestor(userKey(ctx, u.Id))).
		GetAll(ctx, &actions)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}

	for i, key := range keys {
		actions[i].Id = key.IntID()
		actions[i].UserId = u.Id
	}
	return actions
}

func (chal *Challenge) Completed(ctx context.Context, u *User) []*Action {
	return u.getActions(ctx, func(q *datastore.Query) *datastore.Query {
		q = q.Filter("ChallengeId =", chal.Id).Filter("CauseId =", chal.CauseId)
		if chal.Info.EventStart.Time.IsZero() || chal.Info.EventEnd.Time.IsZero() {
			q = q.Filter("When.Time >=", time.Now().Add(-MinChallengeInterval))
		}
		return q
	})
}

func (u *User) Actions(ctx context.Context, after time.Time) []*Action {
	return u.getActions(ctx, func(q *datastore.Query) *datastore.Query {
		if after.IsZero() {
			return q
		}
		return q.Filter("When.Time >=", after)
	})
}
