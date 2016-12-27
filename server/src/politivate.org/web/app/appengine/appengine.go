package appengine

import (
	"net/http"

	"github.com/jtolds/webhelp/whredir"
	"google.golang.org/appengine"

	"politivate.org/web/app"
)

func init() {
	handler := app.RootHandler
	if !appengine.IsDevAppServer() {
		handler = whredir.RequireHost("www.politivate.org",
			whredir.RequireHTTPS(handler))
	}
	http.Handle("/", handler)
}
