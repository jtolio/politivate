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
			c, err := models.GetCause(ctx, causeId.MustGet(ctx))
			if err != nil {
				webhelp.HandleError(w, r, err)
				return
			}
			h.ServeHTTP(w, webhelp.WithContext(
				r, context.WithValue(ctx, causeKey, c)))
		})
}

func getCause(ctx context.Context) (*models.Cause, error) {
	cause, ok := ctx.Value(causeKey).(*models.Cause)
	if !ok {
		return nil, webhelp.ErrInternalServerError.New("no Cause")
	}
	return cause, nil
}

func serveCause(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	cause, err := getCause(ctx)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}
	webhelp.RenderJSON(w, r, cause)
}

func serveCauses(w http.ResponseWriter, r *http.Request) {
	c, err := models.GetCauses(webhelp.Context(r))
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}
	webhelp.RenderJSON(w, r, c)
}
