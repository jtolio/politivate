package views

var _ = T.MustParse(`{{ template "header" (makepair . "New Challenge") }}

{{ template "footer" . }}`)
