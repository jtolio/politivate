package views

var _ = T.MustParse(`{{ template "header" (makepair . "About") }}

<br/><br/><br/>
<p class="text-center"><img src="/static/images/under-construction.gif" /></p>

{{ template "footer" . }}`)
