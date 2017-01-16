package models

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/whfatal"
)

type ChallengeRestriction struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Challenge struct {
	Id      int64     `json:"id" datastore:"-"`
	CauseId int64     `json:"cause_id" datastore:"-"`
	Posted  time.Time `json:"posted_ts"`

	Title         string                 `json:"title"`
	Description   string                 `json:"description"`
	Points        int                    `json:"points"`
	Type          string                 `json:"type"`
	Restrictions  []ChallengeRestriction `json:"restrictions"`
	Deadline      time.Time              `json:"deadline,omitempty"`
	Start         time.Time              `json:"start,omitempty"`
	Database      string                 `json:"database"`
	DirectPhone   string                 `json:"direct_phone"`
	DirectAddress string                 `json:"direct_addr"`
}

func (cause *Cause) NewChallenge(ctx context.Context) *Challenge {
	if cause.Id == 0 {
		whfatal.Error(Error.New("must create Cause first"))
	}

	return &Challenge{
		CauseId: cause.Id,
		Posted:  time.Now(),
	}
}

func challengeKey(ctx context.Context, id, causeId int64) *datastore.Key {
	if causeId == 0 {
		whfatal.Error(Error.New("must create cause first"))
	}
	return datastore.NewKey(
		ctx, "Challenge", "", id, causeKey(ctx, causeId))
}

func (c *Challenge) Save(ctx context.Context) {
	k, err := datastore.Put(ctx, challengeKey(ctx, c.Id, c.CauseId), c)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	c.Id = k.IntID()
}

func (cause *Cause) GetChallenge(ctx context.Context, id int64) *Challenge {
	challenge := Challenge{}
	err := datastore.Get(ctx, challengeKey(ctx, id, cause.Id), &challenge)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	challenge.Id = id
	challenge.CauseId = cause.Id
	return &challenge
}

func getChallenges(ctx context.Context, causeId int64) []*Challenge {
	challenges := make([]*Challenge, 0) // so the json doesn't look like `null`
	if causeId == 0 {
		return challenges
	}
	keys, err := datastore.NewQuery("Challenge").
		Ancestor(causeKey(ctx, causeId)).GetAll(ctx, &challenges)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	for i, key := range keys {
		challenges[i].Id = key.IntID()
		challenges[i].CauseId = causeId
	}
	return challenges
}

func (cause *Cause) GetChallenges(ctx context.Context) []*Challenge {
	return getChallenges(ctx, cause.Id)
}

func GetChallenges(ctx context.Context, causeIds ...int64) []*Challenge {
	challenges := make([]*Challenge, 0) // so the json doesn't look like `null`
	for _, causeId := range causeIds {
		challenges = append(challenges, getChallenges(ctx, causeId)...)
	}
	return challenges
}
