package forms

import (
	"bytes"
	"html/template"

	"gopkg.in/webhelp.v1/whfatal"

	"politivate.org/web/views"
)

type Form struct {
	Error    string
	Template string
	Form     map[string]interface{}
	Page     *views.Page
	Action   string
}

func (f *Form) Render(p *views.Page, action string) template.HTML {
	var buf bytes.Buffer
	f.Page = p
	f.Action = action
	err := views.T.Lookup(f.Template).Execute(&buf, f)
	if err != nil {
		whfatal.Error(err)
	}
	return template.HTML(buf.String())
}
