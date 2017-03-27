package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "Get the App!" "Selected" "get") }}

<div class="row"><div class="col-sm-10 col-sm-offset-1">

<h1>Get the App!</h1>

<p>The Politivate app will be coming soon to an iOS or Android app store near
  you, but in the meantime, if you have Android, open the link below on
  your phone to install the demo version!</p>

<p><a href="/static/demo.apk"><img src="/static/images/appicon144.png"></a></p>

</div></div>

{{ template "footer" . }}`)
