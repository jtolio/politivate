package appengine

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/app"
)

func init() {
	http.Handle("/", webhelp.RequireHost("www.politivate.org",
		webhelp.RequireHTTPS(app.RootHandler)))
}
