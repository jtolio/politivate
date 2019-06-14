package controllers

import (
	"politivate.org/web/controllers/api"
)

func init() {
	mux["api"] = api.Handler
}
