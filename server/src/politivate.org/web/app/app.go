package app

import (
	"github.com/jtolds/webhelp/whcompat"
	"github.com/jtolds/webhelp/whfatal"
	"github.com/jtolds/webhelp/whlog"
	"github.com/jtolds/webhelp/whsess"

	"politivate.org/web/controllers"
	"politivate.org/web/secrets"
)

var (
	RootHandler = whcompat.DoneNotify(whlog.LogRequests(
		whsess.HandlerWithStore(whsess.NewCookieStore(secrets.CookieSecret),
			whfatal.Catch(controllers.Handler))))
)
