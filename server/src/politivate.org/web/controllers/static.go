package controllers

import (
	"net/http"
)

func init() {
	mux["static"] = http.FileServer(http.Dir("./static/"))
}
