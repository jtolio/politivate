package views

var _ = T.MustParse(`{{ template "header" . }}

<h1>Landing!</h1>

<p><a href="/settings">Settings</a></p>

{{ template "footer" . }}`)
