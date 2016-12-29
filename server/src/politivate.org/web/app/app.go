package app

import (
	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whlog"
	"gopkg.in/webhelp.v1/whsess"

	"politivate.org/web/controllers"
	"politivate.org/web/secrets"
)

var (
	RootHandler = whcompat.DoneNotify(whlog.LogRequests(whlog.Default,
		whsess.HandlerWithStore(whsess.NewCookieStore(secrets.CookieSecret),
			whfatal.Catch(controllers.Handler))))
)
