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
	challengeMux["admin"] = whmux.Dir{
		"edit": whmux.Method{
			"GET":  whmux.ExactPath(http.HandlerFunc(editChallengeForm)),
			"POST": whmux.ExactPath(http.HandlerFunc(editChallengeFormAction)),
		},
	}
}

func editChallengeForm(w http.ResponseWriter, r *http.Request) {
	views.Render(w, r, "edit_challenge", map[string]interface{}{
		"Form": forms.EditChallengeForm(administerChallenge(r)),
	})
}

func editChallengeFormAction(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	c := administerChallenge(r)
	ok, f := forms.ProcessChallengeForm(c, r)
	if !ok {
		views.Render(w, r, "edit_challenge", map[string]interface{}{"Form": f})
		return
	}
	c.Save(ctx)

	whfatal.Redirect(fmt.Sprintf("/cause/%d/challenge/%d", c.CauseId, c.Id))
}
