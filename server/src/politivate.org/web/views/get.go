package views

var _ = T.MustParse(`{{ template "header" . }}

<h1>Get the App!</h1>

<p><img src="/static/images/appicon144.png"></p>

{{ template "footer" . }}`)
