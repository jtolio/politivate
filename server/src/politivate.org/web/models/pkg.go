package models

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/spacemonkeygo/errors"
	"github.com/spacemonkeygo/errors/errhttp"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"gopkg.in/webhelp.v1/whfatal"
)

var (
	Error    = errors.NewClass("model error")
	NotFound = Error.NewClass("not found",
		errhttp.SetStatusCode(404), errors.NoCaptureStack())
)

func wrapErr(err error) error {
	if Error.Contains(err) {
		return err
	}
	if err == datastore.ErrNoSuchEntity {
		return NotFound.Wrap(err)
	}
	return Error.Wrap(err)
}

func token() string {
	var token [32]byte
	_, err := rand.Read(token[:])
	if err != nil {
		whfatal.Error(err)
	}
	return hex.EncodeToString(token[:])
}

func deleteAll(ctx context.Context, q func() *datastore.Query) {
	keys, err := q().KeysOnly().GetAll(ctx, nil)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
	if len(keys) == 0 {
		return
	}
	err = datastore.DeleteMulti(ctx, keys)
	if err != nil {
		whfatal.Error(wrapErr(err))
	}
}
