package auth

import (
	"net/http"
	"net/url"

	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/gplus"
	"github.com/markbates/goth/providers/twitter"
	"golang.org/x/net/context"
	"gopkg.in/go-webhelp/whgoth.v1"
	"gopkg.in/webhelp.v1"
	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whroute"

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
	userKey                = webhelp.GenSym()

	LogoutURL = auth.LogoutURL
)

func getCookieUser(ctx context.Context) *models.User {
	gu, err := auth.User(ctx)
	if err != nil {
		whfatal.Error(err)
	}
	if gu == nil {
		return nil
	}
	return models.FindUser(ctx, gu)
}

func getTokenUser(r *http.Request) *models.User {
	token := r.Header.Get("X-Auth-Token")
	if token == "" {
		return nil
	}
	return models.GetUserByAuthToken(whcompat.Context(r), token)
}

func WebLoginRequired(h http.Handler) http.Handler {
	return whroute.HandlerFunc(h,
		func(w http.ResponseWriter, r *http.Request) {
			ctx := whcompat.Context(r)
			u := getCookieUser(ctx)
			if u == nil {
				whfatal.Redirect(LoginRedirect(r.RequestURI))
			}
			h.ServeHTTP(w, whcompat.WithContext(r,
				context.WithValue(ctx, userKey, u)))
		})
}

func APILoginRequired(h http.Handler) http.Handler {
	return whroute.HandlerFunc(h,
		func(w http.ResponseWriter, r *http.Request) {
			ctx := whcompat.Context(r)
			u := getTokenUser(r)
			if u == nil {
				u = getCookieUser(ctx)
			}
			if u == nil {
				whfatal.Error(wherr.Unauthorized.New("X-Auth-Token required"))
			}
			h.ServeHTTP(w, whcompat.WithContext(r,
				context.WithValue(ctx, userKey, u)))
		})
}

func User(ctx context.Context) *models.User {
	u, ok := ctx.Value(userKey).(*models.User)
	if !ok || u == nil {
		whfatal.Error(wherr.Unauthorized.New("login required"))
	}
	return u
}

func LoginRedirect(redirectTo string) string {
	return "/login?" + url.Values{"redirect_to": {redirectTo}}.Encode()
}
