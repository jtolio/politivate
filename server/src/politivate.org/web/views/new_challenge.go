package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "New Challenge") }}

<h1>Create a new challenge</h1>

{{ .Values.Form.Render . "Create" }}

{{ template "footer" . }}`)
