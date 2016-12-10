package controllers

import (
	"net/http"

	"github.com/jtolds/webhelp"

	"politivate.org/web/views"
)

var (
	Mux = webhelp.DirMux{}
)

func Render(w http.ResponseWriter, r *http.Request, template string,
	values map[string]interface{}) {
	tmpl := views.Templates.Lookup(template)
	if tmpl == nil {
		webhelp.HandleError(w, r, webhelp.ErrInternalServerError.New(
			"no template %#v registered", template))
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err := tmpl.Execute(w, values)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}
}
