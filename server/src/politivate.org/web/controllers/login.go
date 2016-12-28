package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp/whmux"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["login"] = whmux.Exact(http.HandlerFunc(loginPage))
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	providers := map[string]string{}
	for _, provider := range auth.Providers() {
		providers[provider.Name()] = provider.LoginURL(r.FormValue("redirect_to"))
	}
	Render(w, r, "login", map[string]interface{}{
		"Providers": providers})
}
