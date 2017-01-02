package controllers

import (
	"politivate.org/web/controllers/auth"
)

func init() {
	mux["profile"] = auth.WebLoginRequired(simpleHandler("profile"))
}
