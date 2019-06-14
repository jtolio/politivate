package controllers

import (
	"net/http"
	"strconv"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/auth"
)

func init() {
	mux["login"] = whmux.Dir{
		"": whmux.RequireGet(http.HandlerFunc(loginPage)),
		"location": whmux.ExactPath(auth.WebLoginRequiredNoLocation(
			whmux.Method{
				"GET":  http.HandlerFunc(locationPage),
				"POST": http.HandlerFunc(locationSet),
			}))}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	redirectTo := r.FormValue("redirect_to")
	if redirectTo == "" {
		redirectTo = "/"
	}
	redirectTo = auth.LoginLocationURL(redirectTo)
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

func locationPage(w http.ResponseWriter, r *http.Request) {
	redirectTo := r.FormValue("redirect_to")
	if redirectTo == "" {
		redirectTo = "/"
	}
	if auth.User(r).LocationSet {
		whfatal.Redirect(redirectTo)
	}
	Render(w, r, "login_location", nil)
}

func locationSet(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	redirectTo := r.FormValue("redirect_to")
	if redirectTo == "" {
		redirectTo = "/"
	}

	latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.Wrap(err))
	}

	longitude, err := strconv.ParseFloat(r.FormValue("longitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.Wrap(err))
	}

	u := auth.User(r)
	u.LocationSet = true
	u.Latitude = latitude
	u.Longitude = longitude
	u.Save(ctx)

	whfatal.Redirect(redirectTo)
}
