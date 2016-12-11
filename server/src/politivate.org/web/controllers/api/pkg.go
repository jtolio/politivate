package api

import (
	"github.com/jtolds/webhelp"
	"github.com/spacemonkeygo/spacelog"

	"politivate.org/web/controllers/auth"
)

var (
	Mux     = webhelp.DirMux{}
	Handler = webhelp.HandleErrorsWith(webhelp.JSONErrorHandler,
		auth.LoginRequired(Mux))

	logger = spacelog.GetLogger()
)
