package views

var _ = T.MustParse(`{{ template "header" . }}

<h1>Causes</h1>

{{ template "footer" . }}`)
