package models

import (
	"github.com/spacemonkeygo/errors"
	"github.com/spacemonkeygo/errors/errhttp"
)

var (
	NotFound = errors.NewClass("not found",
		errhttp.SetStatusCode(404), errors.NoCaptureStack())
)
