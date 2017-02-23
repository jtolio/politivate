package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whmux"
	"gopkg.in/webhelp.v1/whredir"

	"politivate.org/web/views"
)

var (
	mux = whmux.Dir{
		"favicon.ico": whmux.Exact(whredir.RedirectHandler("/static/favicon.ico")),
	}

	Handler http.Handler = mux

	simpleHandler = views.SimpleHandler
	Render        = views.Render
)
