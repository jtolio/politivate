package controllers

import (
	"politivate.org/web/controllers/auth"
)

func init() {
	mux["auth"] = auth.Handler
}
