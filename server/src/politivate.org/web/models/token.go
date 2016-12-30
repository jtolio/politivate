package models

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
)

type AuthToken struct {
	UserId   int64     `datastore:"-" json:"user_id"`
	Token    string    `json:"token"`
	Creation time.Time `json:"creation"`
}

func (u *User) newAuthToken(ctx context.Context) *AuthToken {
	if u.Id == 0 {
		whfatal.Error(Error.New("incomplete user"))
	}

	var token [32]byte
	_, err := rand.Read(token[:])
	if err != nil {
		whfatal.Error(err)
	}

	// TODO: store auth session information so we can double check that the
	//       user is still good later.
	at := &AuthToken{
		UserId:   u.Id,
		Token:    hex.EncodeToString(token[:]),
		Creation: time.Now(),
	}

	_, err = datastore.Put(ctx,
		datastore.NewKey(ctx, "AuthToken", "", 0, userKey(ctx, u.Id)), at)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}

	return at
}

func GetUserByAuthToken(ctx context.Context, token string) *User {
	var tokens []*AuthToken
	keys, err := datastore.NewQuery("AuthToken").Filter("Token =", token).
		GetAll(ctx, &tokens)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	if len(keys) == 0 {
		return nil
	}

	if len(keys) > 1 {
		whfatal.Error(wherr.InternalServerError.New("multiple auth tokens"))
	}

	uk := keys[0].Parent()
	var u User
	err = datastore.Get(ctx, uk, &u)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	u.Id = uk.IntID()
	return &u
}
