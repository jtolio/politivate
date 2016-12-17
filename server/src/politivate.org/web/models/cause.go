package models

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Cause struct {
	Id      int64  `json:"id" datastore:"-"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
}

func NewCause(ctx context.Context) *Cause {
	return &Cause{}
}

func causeKey(ctx context.Context, id int64) *datastore.Key {
	return datastore.NewKey(ctx, "Cause", "", id, nil)
}

func (c *Cause) Save(ctx context.Context) error {
	k, err := datastore.Put(ctx, causeKey(ctx, c.Id), c)
	if err != nil {
		return wrapErr(err)
	}
	c.Id = k.IntID()
	return nil
}

func GetCause(ctx context.Context, id int64) (*Cause, error) {
	cause := Cause{}
	err := wrapErr(datastore.Get(ctx, causeKey(ctx, id), &cause))
	cause.Id = id
	return &cause, err
}

func GetCauses(ctx context.Context) ([]*Cause, error) {
	var causes []*Cause
	keys, err := datastore.NewQuery("Cause").Order("Name").GetAll(ctx, &causes)
	if err != nil {
		return nil, wrapErr(err)
	}
	for i, key := range keys {
		causes[i].Id = key.IntID()
	}
	return causes, nil
}
