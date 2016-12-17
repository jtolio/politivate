package models

import (
	"time"

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

func (cause *Cause) NewChallenge(ctx context.Context) (*Challenge, error) {
	if cause.Id == 0 {
		return nil, Error.New("must create Cause first")
	}

	return &Challenge{
		CauseId: cause.Id,
		Posted:  time.Now(),
	}, nil
}

func challengeKey(ctx context.Context, id, causeId int64) (
	*datastore.Key, error) {
	if causeId == 0 {
		return nil, Error.New("must create cause first")
	}
	return datastore.NewKey(
		ctx, "Challenge", "", id, causeKey(ctx, causeId)), nil
}

func (c *Challenge) Save(ctx context.Context) error {
	ik, err := challengeKey(ctx, c.Id, c.CauseId)
	if err != nil {
		return err
	}
	k, err := datastore.Put(ctx, ik, c)
	if err != nil {
		return wrapErr(err)
	}
	c.Id = k.IntID()
	return nil
}

func (cause *Cause) GetChallenge(ctx context.Context, id int64) (
	*Challenge, error) {
	challenge := Challenge{}
	k, err := challengeKey(ctx, id, cause.Id)
	if err != nil {
		return nil, err
	}
	err = wrapErr(datastore.Get(ctx, k, &challenge))
	challenge.Id = id
	challenge.CauseId = cause.Id
	return &challenge, err
}

func (cause *Cause) GetChallenges(ctx context.Context) ([]*Challenge, error) {
	if cause.Id == 0 {
		return nil, nil
	}
	var challenges []*Challenge
	keys, err := datastore.NewQuery("Challenge").
		Ancestor(causeKey(ctx, cause.Id)).
		Order("Posted").GetAll(ctx, &challenges)
	if err != nil {
		return nil, wrapErr(err)
	}
	for i, key := range keys {
		challenges[i].Id = key.IntID()
		challenges[i].CauseId = cause.Id
	}
	return challenges, nil
}
