package api

import (
	"net/http"

	"github.com/spacemonkeygo/spacelog"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/auth"
)

var (
	authedMux, unauthedMux = whmux.Dir{}, whmux.Dir{}

	Handler http.Handler = wherr.HandleWith(whjson.ErrHandler,
		whfatal.Catch(whmux.Dir{"v1": whmux.Overlay{
			Default: auth.APILoginRequired(authedMux),
			Overlay: unauthedMux,
		}}))

	logger = spacelog.GetLogger()
)
