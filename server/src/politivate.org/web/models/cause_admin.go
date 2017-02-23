package models

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/whfatal"
)

type causeAdmin struct {
	Id      int64 `json:"-", datastore:"-"`
	CauseId int64 `json:"cause_id" datastore:"-"`
	UserId  int64 `json:"user_id" `
}

func (a *causeAdmin) Save(ctx context.Context) {
	if a.CauseId == 0 || a.UserId == 0 {
		whfatal.Error(Error.New("incomplete cause admin"))
	}

	k, err := datastore.Put(ctx, datastore.NewKey(
		ctx, "causeAdmin", "", a.Id, causeKey(ctx, a.CauseId)), a)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	a.Id = k.IntID()
}

func (u *User) Administrate(ctx context.Context, c *Cause) {
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		if !u.IsAdministrating(ctx, c) {
			(&causeAdmin{UserId: u.Id, CauseId: c.Id}).Save(ctx)
		}
		return nil
	}, nil)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
}

func (u *User) Unadministrate(ctx context.Context, c *Cause) {
	deleteAll(ctx, func() *datastore.Query {
		return datastore.NewQuery("causeAdmin").Ancestor(causeKey(ctx, c.Id)).
			Filter("UserId =", u.Id)
	})
}

func (u *User) IsAdministrating(ctx context.Context, c *Cause) bool {
	if u.Id == 0 {
		whfatal.Error(Error.New("must create User first"))
	}
	if c.Id == 0 {
		whfatal.Error(Error.New("must create Cause first"))
	}

	count, err := datastore.NewQuery("causeAdmin").
		Ancestor(causeKey(ctx, c.Id)).Filter("UserId =", u.Id).Count(ctx)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	return count > 0
}
