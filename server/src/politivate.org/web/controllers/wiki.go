package controllers

import (
	"gopkg.in/webhelp.v1/whauth"
	"gopkg.in/webhelp.v1/whredir"

	"politivate.org/web/controllers/wiki"
	"politivate.org/web/secrets"
)

func init() {
	mux["wiki"] = whauth.RequireBasicAuth(whredir.RequireNextSlash(wiki.Handler),
		"Politivate Wiki", secrets.WikiAuth)
}
