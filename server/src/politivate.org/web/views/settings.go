package views

var _ = T.MustParse(`{{ template "header" . }}

<h1>Settings!</h1>

<p><a href="{{.LogoutURL}}">Logout</a></p>

{{ template "footer" . }}`)
