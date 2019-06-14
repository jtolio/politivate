package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"
	"gopkg.in/webhelp.v1/whredir"

	"politivate.org/web/auth"
)

func init() {
	mux["app"] = whmux.Dir{
		"login": whmux.Exact(http.HandlerFunc(appLogin)),
		"token": auth.WebLoginRequired(whmux.Exact(
			http.HandlerFunc(appToken))),
	}
}

func appToken(w http.ResponseWriter, r *http.Request) {
	u := auth.User(r)
	otp := u.NewOTP(whcompat.Context(r))
	whredir.Redirect(w, r,
		fmt.Sprintf("politivate-org-app://www.politivate.org/api/v1/login/otp/%s",
			otp.Token))
}

func appLogin(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	auth.Logout(ctx, w)
	provList, err := auth.Providers(ctx)
	if err != nil {
		whfatal.Error(err)
	}
	providers := map[string]string{}
	for _, provider := range provList {
		providers[provider.Name()] = provider.LoginURL("/app/token")
	}
	Render(w, r, "login", map[string]interface{}{
		"Providers": providers,
	})
}
