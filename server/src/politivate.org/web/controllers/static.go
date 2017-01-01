package controllers

import (
	"politivate.org/web/controllers/static"
)

func init() {
	mux["static"] = static.Handler
}
