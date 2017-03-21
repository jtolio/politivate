package views

var _ = T.MustParse(`{{ template "header" (makepair . "Challenge") }}

{{ if .Values.IsAdministrating }}
  <!-- TODO: figure out how to get the btn-group to work with internal forms -->
  <form method="post" onsubmit="return confirm('Are you sure you want to delete the challenge?');">
    <input type="hidden" name="action" value="delete">
    <div class="btn-group" role="group">
      <button type="submit" class="btn btn-default">Delete Challenge</button>
    </div>
  </form>
{{ end }}

<h1>{{.Values.Challenge.Info.Title}}</h1>

<div class="row">
  <div class="col-sm-8">
    <p>{{.Values.Challenge.Data.Description | format}}</p>
  </div>
</div>

<p><a href="/cause/{{.Values.Cause.Id}}">Back to Cause</a></p>

{{ template "footer" . }}
`)
