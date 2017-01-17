package views

var _ = T.MustParse(`{{ template "header" (makepair . "Get the App!") }}

<h1>Get the App!</h1>

<p>The Politivate app will be coming soon to an iOS or Android app store near
  you!</p>

<p><img src="/static/images/appicon144.png"></p>

{{ template "footer" . }}`)
