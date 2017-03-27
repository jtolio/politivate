package models

import (
	"encoding/json"
	"reflect"
	"sort"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/whfatal"

	"politivate.org/web/controllers/api/gov"
)

type ChallengeRestriction struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Challenge struct {
	Id      int64
	CauseId int64
	Info    ChallengeHeader
	Data    *ChallengeData
}

func (cause *Cause) NewChallenge(ctx context.Context) *Challenge {
	if cause.Id == 0 {
		whfatal.Error(Error.New("must create Cause first"))
	}

	return &Challenge{
		CauseId: cause.Id,
		Info: ChallengeHeader{
			Posted: TimeNow(),
		},
		Data: &ChallengeData{},
	}
}

func (c *Challenge) JSON() map[string]interface{} {
	vals := map[string]interface{}{
		"id":           c.Id,
		"cause_id":     c.CauseId,
		"posted":       c.Info.Posted,
		"title":        c.Info.Title,
		"type":         c.Info.Type,
		"restrictions": c.Info.Restrictions,
		"event_start":  c.Info.EventStart,
		"event_end":    c.Info.EventEnd,
	}
	if c.Data != nil {
		vals["description"] = c.Data.Description
		vals["database"] = c.Data.Database
		vals["direct_phone"] = c.Data.DirectPhone
		vals["direct_address"] = c.Data.DirectAddress
		vals["direct_latitude"] = c.Data.DirectLatitude
		vals["direct_longitude"] = c.Data.DirectLongitude
		vals["direct_radius"] = c.Data.DirectRadius
	}
	return vals
}

func (c *Challenge) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.JSON())
}

type ChallengeHeader struct {
	Posted Time

	Title string

	// Type can be "phonecall" or "location"
	Type string

	Restrictions []ChallengeRestriction `datastore:",noindex"`

	// If neither EventStart or EventEnd are set, there's no timeframe for the
	// challenge.
	// If just EventEnd is set, the challenge has a deadline.
	// If both EventStart and EventEnd are set, then the challenge has a specific
	// timeframe.
	// It doesn't currently make sense to set EventStart only.
	EventStart Time
	EventEnd   Time

	Enabled bool
}

type ChallengeData struct {
	Description string `datastore:",noindex"`

	// Database can currently be "direct", "us", "ushouse", or "ussenate"
	Database string

	DirectPhone string

	DirectAddress   string
	DirectLatitude  float64
	DirectLongitude float64
	DirectRadius    float64
}

func challengeKey(ctx context.Context, id, causeId int64) *datastore.Key {
	if causeId == 0 {
		whfatal.Error(Error.New("must create cause first"))
	}
	return datastore.NewKey(ctx, "Challenge", "", id, causeKey(ctx, causeId))
}

func challengeDataKey(ctx context.Context, id, causeId int64) *datastore.Key {
	if causeId == 0 {
		whfatal.Error(Error.New("must create cause first"))
	}
	if id == 0 {
		whfatal.Error(Error.New("must create challenge to get id"))
	}
	return datastore.NewKey(ctx, "ChallengeData", "", id, causeKey(ctx, causeId))
}

func (c *Challenge) Save(ctx context.Context) {
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		k, err := datastore.Put(ctx, challengeKey(ctx, c.Id, c.CauseId),
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

func (c *Challenge) Delete(ctx context.Context) {
	// first, remove all actions
	// can't remove all actions in a transaction since they have different
	// ancestor keys.
	deleteAll(ctx, func() *datastore.Query {
		return datastore.NewQuery("Action").Filter("ChallengeId =", c.Id).
			Filter("CauseId =", c.CauseId)
	})

	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		err := datastore.Delete(ctx, challengeKey(ctx, c.Id, c.CauseId))
		if err != nil {
			return err
		}
		return datastore.Delete(ctx, challengeDataKey(ctx, c.Id, c.CauseId))
	}, nil)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
}

func (cause *Cause) GetChallenge(ctx context.Context, id int64) *Challenge {
	challenge := Challenge{Data: &ChallengeData{}}
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		err := datastore.Get(ctx, challengeKey(ctx, id, cause.Id),
			&challenge.Info)
		if err != nil {
			return err
		}
		return datastore.Get(ctx, challengeDataKey(ctx, id, cause.Id),
			challenge.Data)
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
	var challengeHeaders []*ChallengeHeader
	if causeId == 0 {
		// use make so the json doesn't look like `null`
		return make([]*Challenge, 0)
	}
	q := datastore.NewQuery("Challenge").Ancestor(
		causeKey(ctx, causeId))
	if mod != nil {
		q = mod(q)
	}
	keys, err := q.GetAll(ctx, &challengeHeaders)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	challenges := make([]*Challenge, len(keys))
	for i, key := range keys {
		challenges[i] = &Challenge{
			Id:      key.IntID(),
			CauseId: causeId,
			Info:    *challengeHeaders[i],
		}
	}
	return challenges
}

func getLiveChallenges(ctx context.Context, causeId int64) []*Challenge {
	return append(getChallengesHelper(ctx, causeId,
		func(q *datastore.Query) *datastore.Query {
			return q.Filter("Enabled =", true).Order("EventEnd.Time").
				Filter("EventEnd.Time >", time.Now())
		}), getChallengesHelper(ctx, causeId,
		func(q *datastore.Query) *datastore.Query {
			return q.Filter("Enabled =", true).
				Filter("EventEnd.Time =", time.Time{})
		})...)
}

func (cause *Cause) GetLiveChallenges(ctx context.Context) []*Challenge {
	return getLiveChallenges(ctx, cause.Id)
}

func (cause *Cause) GetAllChallenges(ctx context.Context) []*Challenge {
	return getChallengesHelper(ctx, cause.Id, nil)
}

func GetLiveChallenges(ctx context.Context, causeIds ...int64) []*Challenge {
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

func (u *User) isTarget(ctx context.Context, districts []gov.FederalDistrict,
	restrictions []ChallengeRestriction) bool {
	if len(restrictions) == 0 {
		return true
	}
	for _, restriction := range restrictions {
		switch restriction.Type {
		case "state":
			for _, district := range districts {
				if restriction.Value == district.State {
					return true
				}
			}
		default:
			logger.Warnf("unknown restriction type: %s", restriction.Type)
			return true
		}
	}
	return false
}

func (u *User) GetLiveChallenges(ctx context.Context) []*Challenge {
	causeIds := u.CauseIds(ctx)
	challenges := make([]*Challenge, 0)
	districts := gov.FederalDistrictLocateByGPS(ctx, u.Latitude, u.Longitude)
	for _, causeId := range causeIds {
		for _, challenge := range getLiveChallenges(ctx, causeId) {
			if !u.isTarget(ctx, districts, challenge.Info.Restrictions) {
				continue
			}
			challenges = append(challenges, challenge)
		}
	}
	// we could probably be faster and make this some kind of merge operation,
	// since all the individual getLiveChallenges calls are already sorted.
	sort.Sort(challengeSorter(challenges))
	return challenges
}
