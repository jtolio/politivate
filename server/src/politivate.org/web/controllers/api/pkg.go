package api

import (
	"net/http"

	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"
	"github.com/spacemonkeygo/spacelog"

	"politivate.org/web/controllers/auth"
)

var (
	mux                  = whmux.Dir{}
	Handler http.Handler = wherr.HandleWith(whjson.ErrHandler,
		whfatal.Catch(auth.APILoginRequired(whmux.Dir{"v1": mux})))

	logger = spacelog.GetLogger()
)
