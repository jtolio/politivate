package views

var _ = T.MustParse(`{{ template "header" (makepair . "Login") }}

<h1>Login</h1>

{{ with (index .Values.Providers "facebook") }}
  <p><a class="btn btn-primary btn-lg" href="{{.}}"
        role="button">Login with Facebook</a></p>
{{ end }}

{{ with (index .Values.Providers "gplus") }}
  <p><a class="btn btn-primary btn-lg" href="{{.}}"
        role="button">Login with Google</a></p>
{{ end }}

{{ with (index .Values.Providers "twitter") }}
  <p><a class="btn btn-primary btn-lg" href="{{.}}"
        role="button">Login with Twitter</a></p>
{{ end }}

{{ template "footer" . }}`)
