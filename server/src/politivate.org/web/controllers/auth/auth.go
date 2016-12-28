package auth

import (
	"net/http"
	"net/url"

	"github.com/jtolds/webhelp-whgoth"
	"github.com/jtolds/webhelp/wherr"
	"github.com/jtolds/webhelp/whfatal"
	"github.com/jtolds/webhelp/whredir"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/gplus"
	"github.com/markbates/goth/providers/twitter"
	"golang.org/x/net/context"

	"politivate.org/web/models"
	"politivate.org/web/secrets"
)

var (
	auth = whgoth.NewAuthProviders(
		"/auth", "auth",
		gplus.New(
			secrets.GoogleClientId,
			secrets.GoogleClientSecret,
			"https://www.politivate.org/auth/provider/gplus/callback"),
		facebook.New(
			secrets.FacebookClientId,
			secrets.FacebookClientSecret,
			"https://www.politivate.org/auth/provider/facebook/callback"),
		twitter.New(
			secrets.TwitterClientId,
			secrets.TwitterClientSecret,
			"https://www.politivate.org/auth/provider/twitter/callback"),
	)

	Handler   http.Handler = auth
	Providers              = auth.Providers

	LogoutURL = auth.LogoutURL
)

func WebLoginRequired(h http.Handler) http.Handler {
	return auth.RequireUser(h, whredir.RedirectHandlerFunc(
		func(r *http.Request) string {
			return LoginRedirect(r.RequestURI)
		}))
}

func APILoginRequired(h http.Handler) http.Handler {
	return auth.RequireUser(h, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			wherr.Handle(w, r, wherr.Unauthorized.New("login required"))
		}))
}

func User(ctx context.Context) *models.User {
	u, err := auth.User(ctx)
	if err != nil {
		whfatal.Error(err)
	}
	if u == nil {
		return nil
	}
	return models.FindUser(ctx, u)
}

func LoginRedirect(redirectTo string) string {
	return "/login?" + url.Values{"redirect_to": {redirectTo}}.Encode()
}
