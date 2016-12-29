package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/views"
)

var (
	mux                  = whmux.Dir{}
	Handler http.Handler = mux

	Render = views.T.Render
)
