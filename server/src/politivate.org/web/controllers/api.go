package controllers

import (
	"politivate.org/web/controllers/api"
)

func init() {
	Mux["api"] = api.Handler
}
