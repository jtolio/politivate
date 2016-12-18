package cause

import (
	"net/http"

	"github.com/jtolds/webhelp"
	"github.com/spacemonkeygo/spacelog"
	"golang.org/x/net/context"

	"politivate.org/web/models"
)

var (
	causeId  = webhelp.NewIntArgMux()
	causeKey = webhelp.GenSym()

	Handler http.Handler = causeId.Shift(requireCause(mux))
	mux                  = webhelp.DirMux{}

	logger = spacelog.GetLogger()
)

func requireCause(h http.Handler) http.Handler {
	return webhelp.RouteHandlerFunc(h,
		func(w http.ResponseWriter, r *http.Request) {
			ctx := webhelp.Context(r)
			h.ServeHTTP(w, webhelp.WithContext(
				r, context.WithValue(ctx, causeKey,
					models.GetCause(ctx, causeId.MustGet(ctx)))))
		})
}

func mustGetCause(ctx context.Context) *models.Cause {
	return ctx.Value(causeKey).(*models.Cause)
}
