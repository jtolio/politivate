package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/views"
)

var (
	mux                  = webhelp.DirMux{}
	Handler http.Handler = mux

	Render = views.T.Render
)
