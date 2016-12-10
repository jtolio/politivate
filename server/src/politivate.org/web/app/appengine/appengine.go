package appengine

import (
	"net/http"

	"github.com/jtolds/webhelp"
	"github.com/spacemonkeygo/spacelog"

	"politivate.org/web/app"
)

func init() {
	spacelog.SetLevel(nil, spacelog.Debug)
	http.Handle("/", RequireHTTPS("www.politivate.org", app.RootHandler))
}

func RequireHTTPS(host string, handler http.Handler) http.Handler {
	return webhelp.RouteHandlerFunc(handler,
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Scheme != "https" || r.URL.Host != host {
				u := *r.URL
				u.Scheme = "https"
				u.Host = host
				webhelp.Redirect(w, r, u.String())
			} else {
				handler.ServeHTTP(w, r)
			}
		})
}
