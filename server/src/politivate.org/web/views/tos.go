package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "Terms of Service") }}

<h1>Terms of Service</h1>

<p>Terms of service coming soon!</p>

{{ template "footer" . }}`)
