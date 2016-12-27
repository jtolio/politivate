package models

import (
	"github.com/jtolds/webhelp/whfatal"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type User struct {
	Id   int64  `json:"id" datastore:"-"`
	Name string `json:"name"`
}

func NewUser(ctx context.Context) *User {
	return &User{}
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
