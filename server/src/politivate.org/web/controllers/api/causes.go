package api

import (
	"net/http"

	"github.com/jtolds/webhelp"
	"golang.org/x/net/context"

	"politivate.org/web/models"
)

var (
	causeId  = webhelp.NewIntArgMux()
	causeKey = webhelp.GenSym()
	causeMux = webhelp.DirMux{}
)

func init() {
	mux["cause"] = causeId.Shift(requireCause(causeMux))
	mux["causes"] = webhelp.Exact(http.HandlerFunc(serveCauses))
	causeMux[""] = webhelp.RequireGet(http.HandlerFunc(serveCause))
}

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

func serveCause(w http.ResponseWriter, r *http.Request) {
	webhelp.RenderJSON(w, r, mustGetCause(webhelp.Context(r)))
}

func serveCauses(w http.ResponseWriter, r *http.Request) {
	webhelp.RenderJSON(w, r, models.GetCauses(webhelp.Context(r)))
}
