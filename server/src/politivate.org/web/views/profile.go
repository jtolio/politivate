package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "Profile") }}

{{ $user := .Values.User }}

<p><img src="{{$user.AvatarURL}}" width=48 /></p>

<dl>
<dt>Name</dt><dd>{{$user.Name}}</dd>
<dt>NickName</dt><dd>{{$user.NickName}}</dd>
<dt>Email</dt><dd>{{$user.Email}}</dd>
<dt>Latitude</dt><dd>{{$user.Latitude}}</dd>
<dt>Longitude</dt><dd>{{$user.Longitude}}</dd>
</dl>

<ul>
{{ $ctx := .Ctx }}
{{ range .Values.Actions }}
  <li>Completed challenge <b>{{(.Challenge $ctx).Info.Title}}</b> on {{.When.Time.Format "2006/01/02"}}</li>
{{ end }}
</ul>

<p><a href="{{.LogoutURL}}">Logout</a></p>

{{ template "footer" . }}`)
