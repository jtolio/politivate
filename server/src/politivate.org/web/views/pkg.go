package views

import (
	"fmt"
	"html/template"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spacemonkeygo/spacelog"
)

var (
	logger = spacelog.GetLogger()
)

type Pair struct {
	First, Second interface{}
}

var (
	Templates = template.New("_").Funcs(
		template.FuncMap{
			"makepair": func(first, second interface{}) Pair {
				return Pair{First: first, Second: second}
			},
			"safeurl": func(val string) template.URL {
				return template.URL(val)
			}})
)

func register(tmpl string) bool {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to determine template name")
	}
	name := strings.TrimSuffix(filepath.Base(filename), ".go")
	if Templates.Lookup(name) != nil {
		panic(fmt.Sprintf("template %#v already registered", name))
	}

	template.Must(Templates.New(name).Parse(tmpl))

	return true
}
