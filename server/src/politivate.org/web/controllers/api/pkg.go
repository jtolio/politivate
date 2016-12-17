package api

import (
	"net/http"

	"github.com/jtolds/webhelp"
	"github.com/spacemonkeygo/spacelog"
)

var (
	mux                  = webhelp.DirMux{}
	Handler http.Handler = webhelp.HandleErrorsWith(webhelp.JSONErrorHandler,
		webhelp.FatalHandler(webhelp.DirMux{"v1": mux}))

	logger = spacelog.GetLogger()
)
