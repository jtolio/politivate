package controllers

import (
	"github.com/jtolds/webhelp"

	"politivate.org/web/views"
)

var (
	mux     = webhelp.DirMux{}
	Handler = mux

	Render = views.T.Render
)
