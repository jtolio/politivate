package views

var _ = T.MustParse(`{{ template "header" (makepair . "Challenge") }}

{{ if .Values.IsAdministrating }}
  <ul>
    <li><form method="post"><input type="hidden" name="action" value="delete"><input type="submit" value="Delete Challenge"></form></li>
  </ul>
{{ end }}


<h1>{{.Values.Challenge.Info.Title}}</h1>

<p>{{.Values.Challenge.Data.Description}}</p>

<p><a href="/cause/{{.Values.Cause.Id}}">Back to Cause</a></p>

{{ template "footer" . }}
`)
