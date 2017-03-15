package cause

import (
	"fmt"
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/auth"
	"politivate.org/web/models"
	"politivate.org/web/views"
)

var (
	inviteToken = whmux.NewStringArg()
)

func init() {
	mux["admin"] = whmux.Dir{
		"invite": inviteToken.ShiftOpt(
			auth.WebLoginRequired(whmux.Exact(http.HandlerFunc(useInvite))),
			whmux.Exact(http.HandlerFunc(newInvite))),
	}
}

func newInvite(w http.ResponseWriter, r *http.Request) {
	c := administerCause(r)
	views.Render(w, r, "new_invite", map[string]interface{}{
		"Cause":  c,
		"Invite": c.CreateAdminInvite(whcompat.Context(r)),
	})
}

func useInvite(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	c := models.GetCause(ctx, causeId.MustGet(ctx))
	c.UseAdminInvite(ctx, inviteToken.Get(ctx), auth.User(r))
	whfatal.Redirect(fmt.Sprintf("/cause/%d", c.Id))
}
