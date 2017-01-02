package views

var _ = T.MustParse(`{{ template "header" (makepair . "Login") }}

<p><a class="btn btn-primary btn-lg" href="{{index .Values.Providers "facebook"}}" role="button">Login with Facebook</a></p>
<p><a class="btn btn-primary btn-lg" href="{{index .Values.Providers "gplus"}}" role="button">Login with Google</a></p>
<p><a class="btn btn-primary btn-lg" href="{{index .Values.Providers "twitter"}}" role="button">Login with Twitter</a></p>

{{ template "footer" . }}`)
