package api

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/models"
)

func init() {
	unauthedMux["login"] = whmux.Exact(http.HandlerFunc(login))
}

func login(w http.ResponseWriter, r *http.Request) {
	whjson.Render(w, r, models.CreateAuthTokenByOTP(
		whcompat.Context(r), r.FormValue("otp")))
}
