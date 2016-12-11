package auth

import (
	"net/http"
	"net/url"

	"github.com/jtolds/webhelp-oauth2"

	"politivate.org/web/secrets"
)

var (
	Auth = func() *oauth2.ProviderGroup {
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

	Handler = Auth
)

func LoginRequired(h http.Handler) http.Handler {
	return Auth.LoginRequired(h, LoginRedirect)
}

func LoginRedirect(redirectTo string) string {
	return "/login?" + url.Values{"redirect_to": {redirectTo}}.Encode()
}
