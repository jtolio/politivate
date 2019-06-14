package controllers

import (
	"politivate.org/web/auth"
)

func init() {
	mux["auth"] = auth.Handler
}
