package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "Privacy policy") }}

<h1>Privacy policy</h1>

<p>Privacy policy coming soon!</p>

{{ template "footer" . }}`)
