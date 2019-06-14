package auth

import (
	"net/http"
	"net/url"

	"golang.org/x/net/context"
	"gopkg.in/go-webhelp/whgoth.v1"
	"gopkg.in/webhelp.v1"
	"gopkg.in/webhelp.v1/whcache"
	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whroute"

	"politivate.org/web/models"
	"politivate.org/web/secrets"
)

var (
	auth = whgoth.NewLazyAuthProviders("/auth", "auth",
		secrets.Providers)
	Handler   http.Handler = auth
	Providers              = auth.Providers
	userKey                = webhelp.GenSym()
	LogoutURL              = auth.LogoutURL
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

func User(r *http.Request) *models.User {
	ctx := whcompat.Context(r)
	if u, ok := whcache.Get(ctx, userKey).(*models.User); ok {
		return u
	}
	u := getTokenUser(r)
	if u == nil {
		u = getCookieUser(ctx)
	}
	whcache.Set(ctx, userKey, u)
	return u
}

func WebLoginRequired(h http.Handler) http.Handler {
	return whroute.HandlerFunc(h,
		func(w http.ResponseWriter, r *http.Request) {
			u := User(r)
			if u == nil {
				whfatal.Redirect(LoginURL(r.RequestURI))
			}
			if !u.LocationSet {
				whfatal.Redirect(LoginLocationURL(r.RequestURI))
			}
			h.ServeHTTP(w, r)
		})
}

func WebLoginRequiredNoLocation(h http.Handler) http.Handler {
	return whroute.HandlerFunc(h,
		func(w http.ResponseWriter, r *http.Request) {
			if User(r) == nil {
				whfatal.Redirect(LoginURL(r.RequestURI))
			}
			h.ServeHTTP(w, r)
		})
}

func APILoginRequired(h http.Handler) http.Handler {
	return whroute.HandlerFunc(h,
		func(w http.ResponseWriter, r *http.Request) {
			if u := User(r); u == nil || !u.LocationSet {
				whfatal.Error(wherr.Unauthorized.New("X-Auth-Token required"))
			}
			h.ServeHTTP(w, r)
		})
}

func LoginURL(redirectTo string) string {
	return "/login?" + url.Values{"redirect_to": {redirectTo}}.Encode()
}

func LoginLocationURL(redirectTo string) string {
	return "/login/location?" +
		url.Values{"redirect_to": {redirectTo}}.Encode()
}

func Logout(ctx context.Context, w http.ResponseWriter) {
	err := auth.Logout(ctx, w)
	if err != nil {
		whfatal.Error(err)
	}
}
