package models

import (
	"github.com/spacemonkeygo/errors"
	"github.com/spacemonkeygo/errors/errhttp"
	"google.golang.org/appengine/datastore"
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
