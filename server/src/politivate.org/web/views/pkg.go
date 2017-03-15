package views

import (
	"html/template"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/context"
	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whmux"
	"gopkg.in/webhelp.v1/whtmpl"

	"politivate.org/web/auth"
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
	Beta   bool
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
	u := auth.User(r)
	T.Render(w, r, template, &Page{
		User:   u,
		Beta:   u != nil && u.BetaTester,
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

var (
	noNewlines        = regexp.MustCompile(`[\t\f\r ]+`)
	oneNewline        = regexp.MustCompile(`[ ]*\n`)
	twoOrMoreNewlines = regexp.MustCompile(`\n\n+`)
	bold1             = regexp.MustCompile(`\*{2}([^\s]([^\*]*[^\s]|))\*{2}`)
	bold2             = regexp.MustCompile(`_{2}([^\s]([^_]*[^\s]|))_{2}`)
	italic1           = regexp.MustCompile(`\*{1}([^\s]([^\*]*[^\s]|))\*{1}`)
	italic2           = regexp.MustCompile(`_{1}([^\s]([^_]*[^\s]|))_{1}`)
	monospace         = regexp.MustCompile(
		"`" + `{1}([^\s]([^\*]*[^\s]|))` + "`" + `{1}`)
	header1  = regexp.MustCompile(`(^|\n)\n*#([^\n]+)\n*(\n|$)`)
	header2  = regexp.MustCompile(`(^|\n)\n*##([^\n]+)\n*(\n|$)`)
	header3  = regexp.MustCompile(`(^|\n)\n*###([^\n]+)\n*(\n|$)`)
	listitem = regexp.MustCompile(`(^|\n)[ ]*\*[ ]+([^\s]([^\n]*[^\s]))[ ]*`)
	listend  = regexp.MustCompile(`</ul>\n*`)
)

// What the heck?
//
// Yeah, this is terrible. Here's why we're here:
//  * Whatever we use for formatting content must render the same in both
//    React Native and on web pages.
//  * People are starting to become most familiar with Markdown, but we don't
//    need all of Markdown.
//
// So, here is a small collection of lines that implements Markdown-like
// behavior that can easily be ported to Javascript.
//
// Better would be a Go library that parses user-input, supports a reasonably
// small subset of Markdown (but not tables? maybe?), and returns a JSONifiable
// structure that we can turn into HTML or return to the app. I haven't found
// anything like that yet; we'll probably have to write it.
// TODO: ^
func format(data string) template.HTML {
	data = strings.TrimSpace(data)
	data = noNewlines.ReplaceAllString(data, " ")
	data = oneNewline.ReplaceAllString(data, "\n")
	data = twoOrMoreNewlines.ReplaceAllString(data, "\n\n")

	data = template.HTMLEscapeString(data)

	data = header3.ReplaceAllString(data, `<h4>$2</h4>`)
	data = header2.ReplaceAllString(data, `<h3>$2</h3>`)
	data = header1.ReplaceAllString(data, `<h2>$2</h2>`)
	data = listitem.ReplaceAllString(data, `<ul><li>$2</li></ul>`)
	data = strings.Replace(data, `</ul><ul>`, "", -1)
	data = listend.ReplaceAllString(data, `</ul>`)
	data = strings.Replace(data, "\n\n",
		`<span class="small-break"></span>`, -1)
	data = strings.Replace(data, "\n", "<br/>", -1)
	data = bold1.ReplaceAllString(data, `<strong>$1</strong>`)
	data = bold2.ReplaceAllString(data, `<strong>$1</strong>`)
	data = italic1.ReplaceAllString(data, `<em>$1</em>`)
	data = italic2.ReplaceAllString(data, `<em>$1</em>`)
	data = monospace.ReplaceAllString(data, `<code>$1</code>`)
	return template.HTML(data)
}
