package auth

import (
	"net/http"
	"net/url"

	"github.com/jtolds/webhelp-oauth2"
	"golang.org/x/net/context"

	"politivate.org/web/models"
	"politivate.org/web/secrets"
)

var (
	auth = func() *oauth2.ProviderGroup {
		group, err := oauth2.NewProviderGroup(
			"auth", "/auth", oauth2.RedirectURLs{},
			oauth2.Google(oauth2.Config{
				ClientID:     secrets.GoogleClientId,
				ClientSecret: secrets.GoogleClientSecret,
				Scopes:       []string{"profile", "email"},
				RedirectURL:  "https://www.politivate.org/auth/google/_cb"}),
			oauth2.Facebook(oauth2.Config{
				ClientID:     secrets.FacebookClientId,
				ClientSecret: secrets.FacebookClientSecret,
				RedirectURL:  "https://www.politivate.org/auth/facebook/_cb"}))
		if err != nil {
			panic(err)
		}
		return group
	}()

	Handler   http.Handler = auth
	Providers              = auth.Providers
	LogoutURL              = auth.LogoutAllURL
)

func WebLoginRequired(h http.Handler) http.Handler {
	// TODO
	// 	return Auth.LoginRequired(h, LoginRedirect)
	return h
}

func APILoginRequired(h http.Handler) http.Handler {
	// TODO
	// Should return an error if the user isn't logged in
	return h
}

func User(ctx context.Context) *models.User {
	// TODO
	return models.GetUsers(ctx)[0]
}

func LoginRedirect(redirectTo string) string {
	return "/login?" + url.Values{"redirect_to": {redirectTo}}.Encode()
}
