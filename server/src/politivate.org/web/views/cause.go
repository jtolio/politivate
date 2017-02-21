package views

var _ = T.MustParse(`{{ template "header" (makepair . .Values.Cause.Info.Name) }}

{{ if .Values.IsAdministrating }}
  <p><a href="/cause/{{ .Values.Cause.Id }}/challenges/new"
      >New Challenge</a></p>
{{ end }}

<h1><img width=24 height=24 src="{{ .Values.Cause.Info.IconURL }}"
    /> {{ .Values.Cause.Info.Name }}</h1>

<p>{{ .Values.Cause.Data.Description }}</p>

<p><a href="{{ .Values.Cause.Info.URL }}">{{ .Values.Cause.Info.URL }}</a></p>

{{ template "footer" . }}`)
