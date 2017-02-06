package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["login"] = whmux.Exact(http.HandlerFunc(loginPage))
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	redirectTo := r.FormValue("redirect_to")
	if redirectTo == "" {
		redirectTo = "/"
	}
	if auth.User(r) != nil {
		whfatal.Redirect(redirectTo)
	}
	provList, err := auth.Providers(whcompat.Context(r))
	if err != nil {
		whfatal.Error(err)
	}
	providers := map[string]string{}
	for _, provider := range provList {
		providers[provider.Name()] = provider.LoginURL(redirectTo)
	}
	Render(w, r, "login", map[string]interface{}{
		"Providers": providers,
	})
}
