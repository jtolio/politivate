package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" .Values.Callenge.Info.Title) }}

<div class="row">
  <div class="col-sm-8">
    <h1 style="margin-bottom: 10px;">{{.Values.Challenge.Info.Title}}</h1>
    <p><a href="/cause/{{.Values.Cause.Id}}">Back to {{ .Values.Cause.Info.Name }}</a></p>

    <p>{{.Values.Challenge.Data.Description | format}}</p>
  </div>
  <div class="col-sm-4">
    <img src="{{ .Values.Cause.Info.IconURL }}" class="img-responsive img-rounded center-block" style="width:100%; margin-bottom:20px;"/>

    {{ if .Values.IsAdministrating }}
      <div class="list-group">
        <a class="list-group-item" href="javascript:doAction('delete', 'Are you sure you want to delete the challenge?');">Delete</a>
      </div>
    {{ end }}
  </div>
</div>

{{ template "footer" . }}`)
