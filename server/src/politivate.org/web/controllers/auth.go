package controllers

import (
	"net/http"
	"net/url"

	"github.com/jtolds/webhelp"
	"github.com/jtolds/webhelp-oauth2"

	"politivate.org/web/secrets"
)

func init() {
	Mux["login"] = webhelp.Exact(http.HandlerFunc(loginPage))
	Mux["auth"] = Auth
}

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
)

func LoginRequired(h http.Handler) http.Handler {
	return Auth.LoginRequired(h, LoginRedirect)
}

func LoginRedirect(redirectTo string) string {
	return "/login?" + url.Values{"redirect_to": {redirectTo}}.Encode()
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	providers := map[string]string{}
	for name, provider := range Auth.Providers() {
		providers[name] = provider.LoginURL(r.FormValue("redirect_to"), false)
	}
	Render(w, r, "login", map[string]interface{}{
		"Providers": providers})
}
