package models

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/whfatal"
)

type CauseAdminInvite struct {
	Token    string    `json:"token"`
	Creation time.Time `json:"creation"`
}

func (c *Cause) CreateAdminInvite(ctx context.Context) *CauseAdminInvite {
	invite := &CauseAdminInvite{
		Token:    token(),
		Creation: time.Now(),
	}
	_, err := datastore.Put(ctx, datastore.NewKey(
		ctx, "CauseAdminInvite", "", 0, causeKey(ctx, c.Id)), invite)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	return invite
}

func (c *Cause) UseAdminInvite(ctx context.Context, token string, u *User) {
	// TODO: constant-time compare somehow
	keys, err := datastore.NewQuery("CauseAdminInvite").
		Ancestor(causeKey(ctx, c.Id)).Filter("Token =", token).KeysOnly().
		GetAll(ctx, nil)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	valid := len(keys) == 1
	err = datastore.DeleteMulti(ctx, keys)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}

	if valid {
		u.Administrate(ctx, c)
	}
}
