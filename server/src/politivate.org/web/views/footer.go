package views

var _ = T.MustParse(`
  {{ template "footerscripts" . }}
  {{ template "footerdoc" . }}
`)
