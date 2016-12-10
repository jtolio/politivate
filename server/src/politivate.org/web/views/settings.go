package views

var _ = register(`{{ template "header" . }}

<h1>Settings!</h1>

<p><a href="{{.LogoutURL}}">Logout</a></p>

{{ template "footer" . }}`)
