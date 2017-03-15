package controllers

import (
	"politivate.org/web/auth"
)

func init() {
	mux["profile"] = auth.WebLoginRequired(Beta(simpleHandler("profile")))
}
