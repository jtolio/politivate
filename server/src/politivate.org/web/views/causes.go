package views

var _ = T.MustParse(`{{ template "header" (makepair . "Causes") }}

<h1>Causes</h1>

{{ template "footer" . }}`)
