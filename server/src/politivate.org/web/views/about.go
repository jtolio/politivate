package views

var _ = T.MustParse(`{{ template "header" . }}

<h1>About</h1>

{{ template "footer" . }}`)
