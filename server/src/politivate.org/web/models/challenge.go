package models

import (
	"reflect"
	"sort"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/whfatal"
)

type NullableTime struct {
	Time time.Time
}

func (t NullableTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return t.Time.MarshalJSON()
}

type ChallengeRestriction struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Challenge struct {
	Id      int64          `json:"id"`
	CauseId int64          `json:"cause_id"`
	Info    ChallengeInfo  `json:"info"`
	Data    *ChallengeData `json:"data,omitempty"`
}

type ChallengeInfo struct {
	Posted NullableTime `json:"posted"`

	Title  string `json:"title"`
	Points int    `json:"points"`

	// Type can be "phonecall" or "location"
	Type string `json:"type"`

	Restrictions []ChallengeRestriction `json:"restrictions"`

	// If neither EventStart or EventEnd are set, there's no timeframe for the
	// challenge.
	// If just EventEnd is set, the challenge has a deadline.
	// If both EventStart and EventEnd are set, then the challenge has a specific
	// timeframe.
	// It doesn't currently make sense to set EventStart only.
	EventStart NullableTime `json:"event_start"`
	EventEnd   NullableTime `json:"event_end"`
}

type ChallengeData struct {
	Description string `json:"description"`

	// Database can currently be "direct", "us", "ushouse", or "ussenate"
	Database string `json:"database"`

	DirectPhone   string `json:"direct_phone"`
	DirectAddress string `json:"direct_addr"`
}

func (cause *Cause) NewChallenge(ctx context.Context) *Challenge {
	if cause.Id == 0 {
		whfatal.Error(Error.New("must create Cause first"))
	}

	return &Challenge{
		CauseId: cause.Id,
		Info: ChallengeInfo{
			Posted: NullableTime{Time: time.Now()},
		},
		Data: &ChallengeData{},
	}
}

func challengeInfoKey(ctx context.Context, id, causeId int64) *datastore.Key {
	if causeId == 0 {
		whfatal.Error(Error.New("must create cause first"))
	}
	return datastore.NewKey(
		ctx, "ChallengeInfo", "", id, causeKey(ctx, causeId))
}

func challengeDataKey(ctx context.Context, id, causeId int64) *datastore.Key {
	if causeId == 0 {
		whfatal.Error(Error.New("must create cause first"))
	}
	if id == 0 {
		whfatal.Error(Error.New("must create challenge info to get id"))
	}
	return datastore.NewKey(
		ctx, "ChallengeData", "", id, causeKey(ctx, causeId))
}

func (c *Challenge) Save(ctx context.Context) {
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		k, err := datastore.Put(ctx, challengeInfoKey(ctx, c.Id, c.CauseId),
			&c.Info)
		if err != nil {
			return err
		}
		c.Id = k.IntID()
		if c.Data != nil {
			_, err = datastore.Put(ctx, challengeDataKey(ctx, c.Id, c.CauseId),
				c.Data)
			if err != nil {
				return err
			}
		}
		return nil
	}, nil)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
}

func (cause *Cause) GetChallenge(ctx context.Context, id int64) *Challenge {
	challenge := Challenge{Data: &ChallengeData{}}
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		err := datastore.Get(ctx, challengeInfoKey(ctx, id, cause.Id),
			&challenge.Info)
		if err != nil {
			return err
		}
		err = datastore.Get(ctx, challengeDataKey(ctx, id, cause.Id),
			challenge.Data)
		if err != nil {
			return err
		}
		return nil
	}, nil)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	challenge.Id = id
	challenge.CauseId = cause.Id
	return &challenge
}

func structFields(val interface{}) []string {
	typ := reflect.ValueOf(val).Type()
	fields := make([]string, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		fields = append(fields, typ.Field(i).Name)
	}
	return fields
}

func getChallengesHelper(ctx context.Context, causeId int64,
	mod func(q *datastore.Query) *datastore.Query) []*Challenge {
	var challengeInfos []*ChallengeInfo
	if causeId == 0 {
		// use make so the json doesn't look like `null`
		return make([]*Challenge, 0)
	}
	q := datastore.NewQuery("ChallengeInfo").Ancestor(causeKey(ctx, causeId))
	if mod != nil {
		q = mod(q)
	}
	keys, err := q.GetAll(ctx, &challengeInfos)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	challenges := make([]*Challenge, len(keys))
	for i, key := range keys {
		challenges[i] = &Challenge{
			Id:      key.IntID(),
			CauseId: causeId,
			Info:    *challengeInfos[i],
		}
	}
	return challenges
}

func getLiveChallenges(ctx context.Context, causeId int64) []*Challenge {
	return append(getChallengesHelper(ctx, causeId,
		func(q *datastore.Query) *datastore.Query {
			return q.Order("EventEnd.Time").Filter("EventEnd.Time >", time.Now())
		}), getChallengesHelper(ctx, causeId,
		func(q *datastore.Query) *datastore.Query {
			return q.Filter("EventEnd.Time =", time.Time{})
		})...)
}

func (cause *Cause) GetChallenges(ctx context.Context) []*Challenge {
	return getLiveChallenges(ctx, cause.Id)
}

func GetChallenges(ctx context.Context, causeIds ...int64) []*Challenge {
	// use make so the json doesn't look like `null`
	challenges := make([]*Challenge, 0)
	for _, causeId := range causeIds {
		challenges = append(challenges, getLiveChallenges(ctx, causeId)...)
	}
	// we could probably be faster and make this some kind of merge operation,
	// since all the individual getLiveChallenges calls are already sorted.
	sort.Sort(challengeSorter(challenges))
	return challenges
}

type challengeSorter []*Challenge

func (c challengeSorter) Len() int      { return len(c) }
func (c challengeSorter) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c challengeSorter) Less(i, j int) bool {
	if c[i].Info.EventEnd.Time.IsZero() {
		return false
	}
	if c[j].Info.EventEnd.Time.IsZero() {
		return true
	}
	return c[i].Info.EventEnd.Time.Before(c[j].Info.EventEnd.Time)
}
