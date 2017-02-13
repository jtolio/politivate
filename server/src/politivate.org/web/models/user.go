package models

import (
	"fmt"
	"strings"

	"github.com/markbates/goth"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
)

type User struct {
	Id             int64  `json:"id" datastore:"-"`
	AuthId         string `json:"-"`
	Name           string `json:"name"`
	NickName       string `json:"nick_name"`
	Email          string `json:"email"`
	AvatarURL      string `json:"avatar_url"`
	CanCreateCause bool   `json:"-"`

	LocationSet bool    `json:"location_set"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
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
