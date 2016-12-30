package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whmux"
	"gopkg.in/webhelp.v1/whredir"

	"politivate.org/web/controllers/auth"
)

func init() {
	mux["app"] = whmux.Dir{
		"login": whmux.Exact(http.HandlerFunc(appLogin)),
		"token": auth.WebLoginRequired(whmux.Exact(
			http.HandlerFunc(appToken))),
	}
}

func appToken(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(ctx)
	otp := u.NewOTP(ctx)
	whredir.Redirect(w, r,
		fmt.Sprintf("politivate-org-app://www.politivate.org/api/v1/login/otp/%s",
			otp.Token))
}

func appLogin(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	auth.Logout(ctx, w)
	providers := map[string]string{}
	for _, provider := range auth.Providers() {
		providers[provider.Name()] = provider.LoginURL("/app/token")
	}
	Render(w, r, "login", map[string]interface{}{
		"Providers": providers})
}
