package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "Profile") }}

<p><img src="{{.User.AvatarURL}}" width=48 /></p>

<dl>
<dt>Name</dt><dd>{{.User.Name}}</dd>
<dt>NickName</dt><dd>{{.User.NickName}}</dd>
<dt>Email</dt><dd>{{.User.Email}}</dd>
<dt>Latitude</dt><dd>{{.User.Latitude}}</dd>
<dt>Longitude</dt><dd>{{.User.Longitude}}</dd>
</dl>

<p><a href="{{.LogoutURL}}">Logout</a></p>

{{ template "footer" . }}`)
