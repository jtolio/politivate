package controllers

import (
	"net/http"
)

func init() {
	mux["static"] = http.FileServer(http.Dir("./static/"))
}

//go:generate bash -c "cd ../sass && pyscss application.scss > ../static/css/bootstrap-new.css"
//go:generate bash -c "cd ../static/css && if ! [ -e bootstrap.css ] || ! diff -q bootstrap.css bootstrap-new.css; then cp bootstrap{-new,}.css; fi"
//go:generate rm -f ../static/css/bootstrap-new.css
