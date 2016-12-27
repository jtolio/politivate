package api

import (
	"net/http"

	"github.com/jtolds/webhelp/wherr"
	"github.com/jtolds/webhelp/whfatal"
	"github.com/jtolds/webhelp/whjson"
	"github.com/jtolds/webhelp/whmux"
	"github.com/spacemonkeygo/spacelog"

	"politivate.org/web/controllers/auth"
)

var (
	mux                  = whmux.Dir{}
	Handler http.Handler = wherr.HandleWith(whjson.ErrHandler,
		whfatal.Catch(auth.APILoginRequired(whmux.Dir{"v1": mux})))

	logger = spacelog.GetLogger()
)
