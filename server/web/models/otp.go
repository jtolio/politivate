package models

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
)

type OTP struct {
	UserId   int64 `datastore:"-"`
	Token    string
	Creation Time
}

func (u *User) NewOTP(ctx context.Context) *OTP {
	if u.Id == 0 {
		whfatal.Error(Error.New("incomplete user"))
	}

	// TODO: store auth session information so we can double check that the
	//       user is still good later.
	o := &OTP{
		UserId:   u.Id,
		Token:    token(),
		Creation: TimeNow(),
	}

	_, err := datastore.Put(ctx,
		datastore.NewKey(ctx, "OTP", "", 0, userKey(ctx, u.Id)), o)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}

	return o
}

func CreateAuthTokenByOTP(ctx context.Context, otpToken string) *AuthToken {
	// TODO: constant-time compare somehow
	var otps []*OTP
	keys, err := datastore.NewQuery("OTP").Filter("Token =", otpToken).
		GetAll(ctx, &otps)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	var uk *datastore.Key
	if len(keys) == 1 {
		uk = keys[0].Parent()
	}
	err = datastore.DeleteMulti(ctx, keys)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}

	if uk == nil {
		whfatal.Error(wherr.Unauthorized.New("otp token invalid"))
	}

	var u User
	err = datastore.Get(ctx, uk, &u)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	u.Id = uk.IntID()

	return u.newAuthToken(ctx)
}