package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/markbates/goth"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
)

type User struct {
	Id             int64 `datastore:"-"`
	AuthId         string
	Name           string
	NickName       string
	Email          string
	AvatarURL      string
	CanCreateCause bool

	LocationSet bool
	Latitude    float64
	Longitude   float64
}

func (u *User) JSON() map[string]interface{} {
	vals := map[string]interface{}{
		"id":           u.Id,
		"name":         u.Name,
		"nick_name":    u.NickName,
		"email":        u.Email,
		"avatar_url":   u.AvatarURL,
		"location_set": u.LocationSet,
		"latitude":     u.Latitude,
		"longitude":    u.Longitude,
	}
	return vals
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.JSON())
}

func userKey(ctx context.Context, id int64) *datastore.Key {
	return datastore.NewKey(ctx, "User", "", id, nil)
}

func (u *User) Save(ctx context.Context) {
	k, err := datastore.Put(ctx, userKey(ctx, u.Id), u)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	u.Id = k.IntID()
}

func GetUser(ctx context.Context, id int64) *User {
	user := User{}
	err := datastore.Get(ctx, userKey(ctx, id), &user)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	user.Id = id
	return &user
}

func GetUsers(ctx context.Context) []*User {
	var users []*User
	keys, err := datastore.NewQuery("User").GetAll(ctx, &users)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	for i, key := range keys {
		users[i].Id = key.IntID()
	}
	return users
}

func FindUser(ctx context.Context, user *goth.User) *User {
	authid := fmt.Sprintf("%s:%s",
		strings.Replace(user.Provider, ":", "_", -1),
		user.UserID)
	var users []*User
	keys, err := datastore.NewQuery("User").Filter("AuthId =", authid).
		GetAll(ctx, &users)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	if len(users) > 1 {
		whfatal.Error(wherr.InternalServerError.New(
			"more than one user with same id"))
	}
	if len(users) == 1 {
		users[0].Id = keys[0].IntID()
		return users[0]
	}
	u := &User{
		AuthId:    authid,
		Name:      user.Name,
		NickName:  user.NickName,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
	}
	u.Save(ctx)
	return u
}
