package appengine

import (
	"net/http"

	"github.com/jtolds/webhelp/whgls"
	"github.com/jtolds/webhelp/whredir"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"politivate.org/web/app"
)

func init() {
	handler := app.RootHandler
	if !appengine.IsDevAppServer() {
		handler = whredir.RequireHost("www.politivate.org",
			whredir.RequireHTTPS(handler))
	}
	whgls.SetLogOutput(log.Infof)
	http.Handle("/", whgls.Bind(handler))
}
