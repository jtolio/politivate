package views

var _ = T.MustParse(`{{ template "header" (makepair . .Values.Cause.Info.Name) }}

{{ if .Values.IsAdministrating }}
  <ul>
    <li><a href="/cause/{{ .Values.Cause.Id }}/challenges/new">New Challenge</a></li>
    <li><form method="post"><input type="hidden" name="action" value="delete"><input type="submit" value="Delete Cause"></form></li>
    <li><a href="/cause/{{ .Values.Cause.Id }}/admin/invite">Invite admin</a></li>
  </ul>
{{ end }}

<h1><img width=24 height=24 src="{{ .Values.Cause.Info.IconURL }}"
    /> {{ .Values.Cause.Info.Name }}</h1>

<p>{{ .Values.Cause.Data.Description | format }}</p>

<p><a href="{{ .Values.Cause.Info.URL }}">{{ .Values.Cause.Info.URL }}</a></p>

<h2>Challenges</h2>
<ul>
  {{ $cause := .Values.Cause }}
  {{ range .Values.Challenges }}
    <li><a href="/cause/{{$cause.Id}}/challenge/{{.Id}}">{{.Info.Title}}</a></li>
  {{ end }}
  {{ if (eq (len .Values.Challenges) 0) }}
    <li>No challenges</li>
  {{ end }}
</ul>

{{ template "footer" . }}`)
