package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "Edit Cause") }}

<h1>Edit Cause</h1>

{{ .Values.Form.Render . "Edit" }}

{{ template "footer" . }}`)
