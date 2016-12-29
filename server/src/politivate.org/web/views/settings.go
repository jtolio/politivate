package views

var _ = T.MustParse(`{{ template "header" . }}

<h1>Settings!</h1>

<p><img src="{{.User.AvatarURL}}" /></p>

<dl>
<dt>Name</dt><dd>{{.User.Name}}</dd>
<dt>NickName</dt><dd>{{.User.NickName}}</dd>
<dt>Email</dt><dd>{{.User.Email}}</dd>
</dl>

<p><a href="{{.LogoutURL}}">Logout</a></p>

{{ template "footer" . }}`)
