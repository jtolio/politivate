package views

var _ = T.MustParse(`{{ template "header" (makepair . "About") }}

<h1>About</h1>

{{ template "footer" . }}`)
