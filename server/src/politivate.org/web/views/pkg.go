package views

import (
	"html/template"
	"net/http"
	"strings"

	"golang.org/x/net/context"
	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whmux"
	"gopkg.in/webhelp.v1/whtmpl"

	"politivate.org/web/controllers/auth"
	"politivate.org/web/models"
)

var (
	T = makeCollection()

	mapsAPIKey = `AIzaSyDIh-CmiVPYkNzJ0AVC2RcJZk5JJYCpqqA`
)

func makeCollection() *whtmpl.Collection {
	rv := whtmpl.NewCollection()
	rv.Funcs(template.FuncMap{
		"format": format,
	})
	return rv
}

type Page struct {
	User   *models.User
	Values interface{}
	Req    *http.Request
	Ctx    context.Context
}

func (p *Page) LogoutURL() string {
	return auth.LogoutURL("/")
}

func (p *Page) LoginURL() string {
	return auth.LoginURL(p.Req.RequestURI)
}

func Render(w http.ResponseWriter, r *http.Request, template string,
	values interface{}) {
	if values == nil {
		values = map[string]interface{}{}
	}
	T.Render(w, r, template, &Page{
		User:   auth.User(r),
		Values: values,
		Req:    r,
		Ctx:    whcompat.Context(r),
	})
}

func SimpleHandler(template string) http.Handler {
	return whmux.Exact(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			Render(w, r, template, nil)
		}))
}

func format(data string) template.HTML {
	data = template.HTMLEscapeString(data)
	return template.HTML(strings.Replace(data, "\n", "<br/>\n", -1))
}
