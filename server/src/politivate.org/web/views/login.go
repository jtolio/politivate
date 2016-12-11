package views

var _ = T.MustParse(`{{ template "header" . }}

<h1>Login</h1>
<ul>
{{ range $name, $login := .Providers }}
<li><a href="{{$login}}">{{$name}}</a></li>
{{ end }}
</ul>

{{ template "footer" . }}`)