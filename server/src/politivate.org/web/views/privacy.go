package views

var _ = T.MustParse(`{{ template "header" (makepair . "Privacy") }}

<h1>Privacy</h1>

{{ template "footer" . }}`)
