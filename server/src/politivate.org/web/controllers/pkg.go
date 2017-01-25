package controllers

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whmux"
	"gopkg.in/webhelp.v1/whredir"

	"golang.org/x/net/context"
	"politivate.org/web/controllers/auth"
	"politivate.org/web/models"
	"politivate.org/web/views"
)

var (
	mux = whmux.Dir{
		"favicon.ico": whredir.RedirectHandler("/static/favicon.ico"),
	}
	Handler http.Handler = mux
)

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
	views.T.Render(w, r, template, &Page{
		User:   auth.User(r),
		Values: values,
		Req:    r,
		Ctx:    whcompat.Context(r),
	})
}

func simpleHandler(template string) http.Handler {
	return whmux.Exact(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			Render(w, r, template, nil)
		}))
}
