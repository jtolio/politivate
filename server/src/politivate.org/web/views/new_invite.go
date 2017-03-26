package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "Invite an admin") }}

<p>Send this link to a friend (but don't open it yourself!):
<pre>https://www.politivate.org/cause/{{.Values.Cause.Id}}/admin/invite/{{.Values.Invite.Token}}</pre></p>

<p><a href="/cause/{{.Values.Cause.Id}}" class="btn btn-default">OK</a></p>

{{ template "footer" . }}
`)
