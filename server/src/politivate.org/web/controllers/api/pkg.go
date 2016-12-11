package api

import (
	"encoding/json"
	"net/http"

	"github.com/jtolds/webhelp"
	"github.com/spacemonkeygo/spacelog"

	"politivate.org/web/controllers/auth"
)

var (
	Mux     = webhelp.DirMux{}
	Handler = webhelp.HandleErrorsWith(webhelp.JSONErrorHandler,
		auth.LoginRequired(Mux))

	logger = spacelog.GetLogger()
)

func RenderJSON(w http.ResponseWriter, r *http.Request, value interface{}) {
	data, err := json.MarshalIndent(
		map[string]interface{}{"response": value}, "", "  ")
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}
}
