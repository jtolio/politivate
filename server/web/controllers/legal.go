package controllers

import (
	"gopkg.in/webhelp.v1/whmux"
)

func init() {
	mux["legal"] = whmux.Dir{
		"tos":     simpleHandler("tos"),
		"privacy": simpleHandler("privacy"),
	}
}
