package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["login"] = whmux.Exact(http.HandlerFunc(loginPage))
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	providers := map[string]string{}
	redirectTo := r.FormValue("redirect_to")
	if redirectTo == "" {
		redirectTo = "/"
	}
	for _, provider := range auth.Providers() {
		providers[provider.Name()] = provider.LoginURL(redirectTo)
	}
	Render(w, r, "login", map[string]interface{}{
		"Providers": providers})
}
