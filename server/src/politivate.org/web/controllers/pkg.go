package controllers

import (
	"github.com/jtolds/webhelp"

	"politivate.org/web/views"
)

var (
	Mux     = webhelp.DirMux{}
	Handler = Mux

	Render = views.T.Render
)
