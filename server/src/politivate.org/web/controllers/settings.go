package controllers

import (
	"politivate.org/web/controllers/auth"
)

func init() {
	mux["settings"] = auth.WebLoginRequired(simpleHandler("settings"))
}
