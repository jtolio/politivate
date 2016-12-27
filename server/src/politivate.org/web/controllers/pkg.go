package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp/whmux"

	"politivate.org/web/views"
)

var (
	mux                  = whmux.Dir{}
	Handler http.Handler = mux

	Render = views.T.Render
)
