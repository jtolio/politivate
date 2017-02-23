package cause

import (
	"fmt"
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/auth"
	"politivate.org/web/models"
	"politivate.org/web/views"
)

var (
	inviteToken = whmux.NewStringArg()
)

func init() {
	mux["admin"] = whmux.Dir{
		"invite": inviteToken.ShiftOpt(
			whmux.Exact(http.HandlerFunc(useInvite)),
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
	u := auth.User(r)
	c := models.GetCause(ctx, causeId.MustGet(ctx))
	if u != nil {
		c.UseAdminInvite(ctx, inviteToken.Get(ctx), u)
	}
	whfatal.Redirect(fmt.Sprintf("/cause/%d", c.Id))
}
