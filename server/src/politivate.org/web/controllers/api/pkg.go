package api

import (
	"github.com/jtolds/webhelp"
	"github.com/spacemonkeygo/spacelog"
)

var (
	mux     = webhelp.DirMux{}
	Handler = webhelp.HandleErrorsWith(webhelp.JSONErrorHandler,
		webhelp.DirMux{"v1": mux})

	logger = spacelog.GetLogger()
)
