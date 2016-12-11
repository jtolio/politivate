package views

var _ = T.MustParse(`{{ template "header" . }}

<h1>Privacy</h1>

{{ template "footer" . }}`)
