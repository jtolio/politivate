package appengine

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/app"
)

func init() {
	http.Handle("/", webhelp.RequireHTTPS(
		webhelp.RequireHost("www.politivate.org", app.RootHandler)))
}
