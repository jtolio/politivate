package views

var _ = T.MustParse(`{{ template "header" (makepair . .Values.Cause.Name) }}

<h1><img width=24 height=24 src="{{ .Values.Cause.IconURL }}" /> {{ .Values.Cause.Name }}</h1>

<p>{{ .Values.Cause.Description }}</p>

<p><a href="{{ .Values.Cause.URL }}">{{ .Values.Cause.URL }}</a></p>

{{ template "footer" . }}`)
