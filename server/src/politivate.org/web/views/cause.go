package views

var _ = T.MustParse(`{{ template "header" (makepair . .Values.Cause.Info.Name) }}

{{ if .Values.IsAdministrating }}
  <!-- TODO: figure out how to get the btn-group to work with internal forms -->
  <form method="post" onsubmit="return confirm('Are you sure you want to delete the cause?');">
    <input type="hidden" name="action" value="delete">
    <div class="btn-group" role="group">
      <a href="/cause/{{ .Values.Cause.Id }}/challenges/new" class="btn btn-default">New Challenge</a>
      <a href="/cause/{{ .Values.Cause.Id }}/admin/invite" class="btn btn-default">Invite admin</a>
      <button type="submit" class="btn btn-default">Delete Cause</button>
    </div>
  </form>
{{ end }}

<h1><img width=24 height=24 src="{{ .Values.Cause.Info.IconURL }}"
    /> {{ .Values.Cause.Info.Name }}</h1>

<div class="row">
  <div class="col-sm-8">
    <p>{{ .Values.Cause.Data.Description | format }}</p>
  </div>
</div>

<p><a href="{{ .Values.Cause.Info.URL }}">{{ .Values.Cause.Info.URL }}</a></p>

<h2>Challenges</h2>

  {{ $cause := .Values.Cause }}
  {{ range $i, $chal := .Values.Challenges }}
    {{ if (ne $i 0) }}
      <div class="horizontal-line"></div>
    {{ end }}
    <a href="/cause/{{$cause.Id}}/challenge/{{$chal.Id}}" class="large-button">
      <div class="media">
        <div class="media-left" style="font-size: 25px;">
          {{ if (eq $chal.Info.Type "location") }}
            <span class="glyphicon glyphicon-earphone text-secondary"></span>
          {{ end }}
          {{ if (eq $chal.Info.Type "phonecall") }}
            <span class="glyphicon glyphicon-map-marker text-secondary"></span>
          {{ end }}
        </div>
        <div class="media-body">
          <h4 class="media-heading">{{ $chal.Info.Title }}</h4>
        </div>
      </div>
    </a>
  {{ end }}
  {{ if (eq (len .Values.Challenges) 0) }}
    No challenges.
  {{ end }}

{{ template "footer" . }}`)
