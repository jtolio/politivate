package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"
	"gopkg.in/webhelp.v1/whredir"
	"gopkg.in/webhelp.v1/whroute"

	"politivate.org/web/auth"
	"politivate.org/web/views"
)

var (
	mux = whmux.Dir{
		"favicon.ico": whmux.Exact(whredir.RedirectHandler("/static/favicon.ico")),
		"demo": whmux.Exact(whredir.RedirectHandler(
			"https://appetize.io/app/exy765kug5pkf0qx118c0q9a4r")),
	}

	Handler http.Handler = mux

	simpleHandler = views.SimpleHandler
	Render        = views.Render
)

func Beta(h http.Handler) http.Handler {
	return auth.WebLoginRequired(whroute.HandlerFunc(h,
		func(w http.ResponseWriter, r *http.Request) {
			if !auth.User(r).BetaTester {
				whfatal.Redirect("/")
			}
			h.ServeHTTP(w, r)
		}))
}
