package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "Edit Challenge") }}

<h1>Edit challenge</h1>

{{ .Values.Form.Render . "Edit" }}

{{ template "footer" . }}`)
