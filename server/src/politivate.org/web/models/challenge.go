package models

import (
	"time"

	"github.com/jtolds/webhelp"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Challenge struct {
	Id      int64 `json:"id" datastore:"-"`
	CauseId int64 `json:"cause_id" datastore:"-"`

	Title     string    `json:"title"`
	ShortDesc string    `json:"short_desc"`
	Posted    time.Time `json:"posted_ts"`
	Deadline  time.Time `json:"deadline_ts,omitempty"`
	IconURL   string    `json:"icon_url"`
	Points    int       `json:"points"`
}

func (cause *Cause) NewChallenge(ctx context.Context) *Challenge {
	if cause.Id == 0 {
		webhelp.FatalError(Error.New("must create Cause first"))
	}

	return &Challenge{
		CauseId: cause.Id,
		Posted:  time.Now(),
	}
}

func challengeKey(ctx context.Context, id, causeId int64) *datastore.Key {
	if causeId == 0 {
		webhelp.FatalError(Error.New("must create cause first"))
	}
	return datastore.NewKey(
		ctx, "Challenge", "", id, causeKey(ctx, causeId))
}

func (c *Challenge) Save(ctx context.Context) {
	k, err := datastore.Put(ctx, challengeKey(ctx, c.Id, c.CauseId), c)
	if err != nil {
		webhelp.FatalError(wrapErr(err))
	}
	c.Id = k.IntID()
}

func (cause *Cause) GetChallenge(ctx context.Context, id int64) *Challenge {
	challenge := Challenge{}
	err := datastore.Get(ctx, challengeKey(ctx, id, cause.Id), &challenge)
	if err != nil {
		webhelp.FatalError(wrapErr(err))
	}
	challenge.Id = id
	challenge.CauseId = cause.Id
	return &challenge
}

func (cause *Cause) GetChallenges(ctx context.Context) []*Challenge {
	challenges := make([]*Challenge, 0) // so the json doesn't look like `null`
	if cause.Id == 0 {
		return challenges
	}
	keys, err := datastore.NewQuery("Challenge").
		Ancestor(causeKey(ctx, cause.Id)).
		Order("Posted").GetAll(ctx, &challenges)
	if err != nil {
		webhelp.FatalError(wrapErr(err))
	}
	for i, key := range keys {
		challenges[i].Id = key.IntID()
		challenges[i].CauseId = cause.Id
	}
	return challenges
}
