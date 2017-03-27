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
	Action   string
	Form     map[string]string
}

func (f *Form) Render() template.HTML {
	var buf bytes.Buffer
	err := views.T.Lookup(f.Template).Execute(&buf, f)
	if err != nil {
		whfatal.Error(err)
	}
	return template.HTML(buf.String())
}
