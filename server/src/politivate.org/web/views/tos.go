package views

var _ = T.MustParse(`{{ template "header" . }}

<h1>Terms of service</h1>

{{ template "footer" . }}`)
