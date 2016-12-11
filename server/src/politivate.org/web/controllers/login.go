package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["login"] = webhelp.Exact(http.HandlerFunc(loginPage))
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	providers := map[string]string{}
	for name, provider := range auth.Auth.Providers() {
		providers[name] = provider.LoginURL(r.FormValue("redirect_to"), false)
	}
	Render(w, r, "login", map[string]interface{}{
		"Providers": providers})
}
