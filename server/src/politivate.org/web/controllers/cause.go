package controllers

import (
	"politivate.org/web/controllers/cause"
)

func init() {
	mux["cause"] = cause.Handler
}
