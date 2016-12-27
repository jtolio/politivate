package models

import (
	"github.com/jtolds/webhelp/whfatal"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type Cause struct {
	Id          int64  `json:"id" datastore:"-"`
	Name        string `json:"name"`
	IconURL     string `json:"icon_url" datastore:",noindex"`
	URL         string `json:"url" datastore:",noindex"`
	Description string `json:"description" datastore:",noindex"`
}

func NewCause(ctx context.Context) *Cause {
	return &Cause{}
}

func causeKey(ctx context.Context, id int64) *datastore.Key {
	return datastore.NewKey(ctx, "Cause", "", id, nil)
}

func (c *Cause) Save(ctx context.Context) {
	k, err := datastore.Put(ctx, causeKey(ctx, c.Id), c)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	c.Id = k.IntID()
}

func GetCause(ctx context.Context, id int64) *Cause {
	cause := Cause{}
	err := datastore.Get(ctx, causeKey(ctx, id), &cause)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	cause.Id = id
	return &cause
}

func GetCauses(ctx context.Context) []*Cause {
	causes := make([]*Cause, 0) // so the json doesn't look like `null`
	keys, err := datastore.NewQuery("Cause").Order("Name").GetAll(ctx, &causes)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	for i, key := range keys {
		causes[i].Id = key.IntID()
	}
	return causes
}
