package app

import (
	"github.com/jtolds/webhelp"
	"github.com/jtolds/webhelp/sessions"

	"politivate.org/web/controllers"
	"politivate.org/web/secrets"
)

var (
	RootHandler = webhelp.ContextBase(webhelp.LoggingHandler(
		webhelp.FatalHandler(sessions.HandlerWithStore(
			sessions.NewCookieStore(secrets.CookieSecret),
			controllers.Handler))))
)
