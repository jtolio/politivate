package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "New Cause") }}

<h1>Create Cause</h1>

{{ .Values.Form.Render . "Create" }}

{{ template "footer" . }}`)
