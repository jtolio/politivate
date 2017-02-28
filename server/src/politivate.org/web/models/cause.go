package models

import (
	"encoding/json"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/whfatal"
)

type CauseHeader struct {
	Name    string
	IconURL string `datastore:",noindex"`
	URL     string `datastore:",noindex"`
}

type CauseData struct {
	Description string `datastore:",noindex"`
}

type Cause struct {
	Id   int64
	Info CauseHeader
	Data *CauseData
}

func NewCause(ctx context.Context) *Cause {
	return &Cause{Data: &CauseData{}}
}

func (c *Cause) JSON() map[string]interface{} {
	vals := map[string]interface{}{
		"id":       c.Id,
		"name":     c.Info.Name,
		"icon_url": c.Info.IconURL,
		"url":      c.Info.URL,
	}
	if c.Data != nil {
		vals["description"] = c.Data.Description
	}
	return vals
}

func (c *Cause) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.JSON())
}

func causeKey(ctx context.Context, id int64) *datastore.Key {
	return datastore.NewKey(ctx, "Cause", "", id, nil)
}

func causeDataKey(ctx context.Context, id int64) *datastore.Key {
	if id == 0 {
		whfatal.Error(Error.New("must create cause to get id"))
	}
	return datastore.NewKey(ctx, "CauseData", "", id, causeKey(ctx, id))
}

func (c *Cause) Save(ctx context.Context) {
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		k, err := datastore.Put(ctx, causeKey(ctx, c.Id), &c.Info)
		if err != nil {
			return err
		}
		c.Id = k.IntID()
		if c.Data != nil {
			_, err = datastore.Put(ctx, causeDataKey(ctx, c.Id), c.Data)
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

func GetCause(ctx context.Context, id int64) *Cause {
	cause := Cause{Data: &CauseData{}}
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		err := datastore.Get(ctx, causeKey(ctx, id), &cause.Info)
		if err != nil {
			return err
		}
		return datastore.Get(ctx, causeDataKey(ctx, id), cause.Data)
	}, nil)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	cause.Id = id
	return &cause
}

func GetCauses(ctx context.Context) []*Cause {
	var causeHeaders []*CauseHeader
	keys, err := datastore.NewQuery("Cause").Order("Name").GetAll(
		ctx, &causeHeaders)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	causes := make([]*Cause, len(keys))
	for i, key := range keys {
		causes[i] = &Cause{
			Id:   key.IntID(),
			Info: *causeHeaders[i],
		}
	}
	return causes
}

func (c *Cause) Delete(ctx context.Context) {
	// first, remove all followers
	// can't remove all followers in a transaction since they have different
	// ancestor keys.
	deleteAll(ctx, func() *datastore.Query {
		return datastore.NewQuery("userCause").Filter("CauseId =", c.Id)
	})

	// next, remove all actions
	// different transaction here too
	deleteAll(ctx, func() *datastore.Query {
		return datastore.NewQuery("Action").Filter("CauseId =", c.Id)
	})

	// everything remaining has this cause as an ancestor
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		// remove all challenges
		deleteAll(ctx, func() *datastore.Query {
			return datastore.NewQuery("Challenge").Ancestor(causeKey(ctx, c.Id))
		})
		deleteAll(ctx, func() *datastore.Query {
			return datastore.NewQuery("ChallengeData").Ancestor(causeKey(ctx, c.Id))
		})

		// remove all admins
		deleteAll(ctx, func() *datastore.Query {
			return datastore.NewQuery("causeAdmin").Ancestor(causeKey(ctx, c.Id))
		})

		// remove the cause
		err := datastore.Delete(ctx, causeDataKey(ctx, c.Id))
		if err != nil {
			return err
		}
		return datastore.Delete(ctx, causeKey(ctx, c.Id))
	}, nil)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
}
