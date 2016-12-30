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
	mux["app"] = whmux.Dir{"login": auth.WebLoginRequired(whmux.Exact(
		http.HandlerFunc(appLogin)))}
}

func appLogin(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(ctx)
	otp := u.NewOTP(ctx)
	whredir.Redirect(w, r,
		fmt.Sprintf("politivateapp:///v1/login/otp/%s", otp.Token))
}
