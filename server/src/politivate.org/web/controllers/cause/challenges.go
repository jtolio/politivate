package cause

import (
	"fmt"
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/forms"
	"politivate.org/web/views"
)

func init() {
	mux["challenges"] = whmux.Dir{
		"new": whmux.ExactPath(whmux.Method{
			"GET":  http.HandlerFunc(newChallengeForm),
			"POST": http.HandlerFunc(newChallengeCreate),
		}),
	}
}

func newChallengeForm(w http.ResponseWriter, r *http.Request) {
	views.Render(w, r, "new_challenge", map[string]interface{}{
		"Cause": administerCause(r),
		"Form":  forms.NewChallengeForm(),
	})
}

func newChallengeCreate(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	c := administerCause(r)

	chal := c.NewChallenge(ctx)
	ok, f := forms.ProcessChallengeForm(chal, r)
	if !ok {
		views.Render(w, r, "new_challenge", map[string]interface{}{
			"Cause": c,
			"Form":  f,
		})
		return
	}

	chal.Save(ctx)
	whfatal.Redirect(fmt.Sprintf("/cause/%d", c.Id))
}
