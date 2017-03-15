package views

var _ = T.MustParse(`{{ template "header" (makepair . "Get the App!") }}

<h1>Get the App!</h1>

<p>The Politivate app will be coming soon to an iOS or Android app store near
  you, but in the meantime, if you have Android, open the link below on
  your phone to install the demo version!</p>

<p><a href="/static/demo.apk"><img src="/static/images/appicon144.png"></a></p>

{{ template "footer" . }}`)
