package views

var _ = T.MustParse(`{{ template "header" (makepair . "Terms of Service") }}

<h1>Terms of Service</h1>

{{ template "footer" . }}`)
