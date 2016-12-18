package api

import (
	"net/http"

	"github.com/jtolds/webhelp"
	"github.com/spacemonkeygo/spacelog"

	"politivate.org/web/controllers/auth"
)

var (
	mux                  = webhelp.DirMux{}
	Handler http.Handler = webhelp.HandleErrorsWith(webhelp.JSONErrorHandler,
		webhelp.FatalHandler(auth.APILoginRequired(webhelp.DirMux{"v1": mux})))

	logger = spacelog.GetLogger()
)
