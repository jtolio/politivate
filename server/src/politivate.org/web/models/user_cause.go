package models

import (
	"gopkg.in/webhelp.v1/whfatal"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

type userCause struct {
	Id      int64 `json:"-", datastore:"-"`
	UserId  int64 `json:"user_id" datastore:"-"`
	CauseId int64 `json:"cause_id"`
}

func (u *userCause) Save(ctx context.Context) {
	if u.CauseId == 0 || u.UserId == 0 {
		whfatal.Error(Error.New("incomplete user cause"))
	}

	k, err := datastore.Put(ctx,
		datastore.NewKey(ctx, "userCause", "", u.Id, userKey(ctx, u.UserId)), u)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	u.Id = k.IntID()
}

func (u *User) Follow(ctx context.Context, c *Cause) {
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		if !u.IsFollowing(ctx, c) {
			(&userCause{UserId: u.Id, CauseId: c.Id}).Save(ctx)
		}
		return nil
	}, nil)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
}

func (u *User) Unfollow(ctx context.Context, c *Cause) {
	keys, err := datastore.NewQuery("userCause").
		Ancestor(userKey(ctx, u.Id)).Filter("CauseId =", c.Id).
		KeysOnly().GetAll(ctx, nil)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	for _, k := range keys {
		err = datastore.Delete(ctx, k)
		if err != nil {
			whfatal.Error(wrapErr(err))
		}
	}
}

func (u *User) IsFollowing(ctx context.Context, c *Cause) bool {
	if u.Id == 0 {
		whfatal.Error(Error.New("must create User first"))
	}
	if c.Id == 0 {
		whfatal.Error(Error.New("must create Cause first"))
	}

	count, err := datastore.NewQuery("userCause").
		Ancestor(userKey(ctx, u.Id)).Filter("CauseId =", c.Id).Count(ctx)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	return count > 0
}

func (u *User) CauseIds(ctx context.Context) []int64 {
	var userCauses []*userCause
	_, err := datastore.NewQuery("userCause").
		Ancestor(userKey(ctx, u.Id)).GetAll(ctx, &userCauses)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	causes := make([]int64, 0, len(userCauses))
	for _, uc := range userCauses {
		causes = append(causes, uc.CauseId)
	}
	return causes
}

func (c *Cause) UserCount(ctx context.Context) int64 {
	count, err := datastore.NewQuery("userCause").
		Filter("CauseId =", c.Id).Count(ctx)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	return int64(count)
}
