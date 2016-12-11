package api

import (
	"github.com/jtolds/webhelp"
	"github.com/spacemonkeygo/spacelog"

	"politivate.org/web/controllers/auth"
)

var (
	mux     = webhelp.DirMux{}
	Handler = webhelp.HandleErrorsWith(webhelp.JSONErrorHandler,
		auth.LoginRequired(webhelp.DirMux{"v1": mux}))

	logger = spacelog.GetLogger()
)
