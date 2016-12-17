package appengine

import (
	"net/http"

	"github.com/jtolds/webhelp"
	"google.golang.org/appengine"

	"politivate.org/web/app"
)

func init() {
	handler := app.RootHandler
	if !appengine.IsDevAppServer() {
		handler = webhelp.RequireHost("www.politivate.org",
			webhelp.RequireHTTPS(handler))
	}
	http.Handle("/", handler)
}
