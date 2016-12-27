package cause

import (
	"net/http"

	"github.com/jtolds/webhelp"
	"github.com/jtolds/webhelp/whcompat"
	"github.com/jtolds/webhelp/whmux"
	"github.com/jtolds/webhelp/whroute"
	"github.com/spacemonkeygo/spacelog"
	"golang.org/x/net/context"

	"politivate.org/web/models"
)

var (
	causeId  = whmux.NewIntArg()
	causeKey = webhelp.GenSym()

	Handler http.Handler = causeId.Shift(requireCause(mux))
	mux                  = whmux.Dir{}

	logger = spacelog.GetLogger()
)

func requireCause(h http.Handler) http.Handler {
	return whroute.HandlerFunc(h,
		func(w http.ResponseWriter, r *http.Request) {
			ctx := whcompat.Context(r)
			h.ServeHTTP(w, whcompat.WithContext(
				r, context.WithValue(ctx, causeKey,
					models.GetCause(ctx, causeId.MustGet(ctx)))))
		})
}

func mustGetCause(ctx context.Context) *models.Cause {
	return ctx.Value(causeKey).(*models.Cause)
}
