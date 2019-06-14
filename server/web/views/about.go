package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "About" "Selected" "about") }}

<br/><br/><br/>
<p class="text-center"><img src="/static/images/under-construction.gif" /></p>

{{ template "footer" . }}`)
