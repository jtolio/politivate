package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" .Values.Cause.Info.Name) }}

<div class="row">
  <div class="col-sm-8">
    <h1 style="margin-bottom: 10px;">{{ .Values.Cause.Info.Name }}</h1>
    {{if (ne (len .Values.Cause.Data.ShortDescription) 0) }}
      <p>{{ .Values.Cause.Data.ShortDescription | format }}</p>
    {{end}}

    <ul class="nav nav-tabs" role="tablist">
      <li role="presentation" class="active"><a href="#challenges"
          aria-controls="challenges" role="tab" data-toggle="tab">Challenges</a></li>
      <li role="presentation"><a href="#about"
          aria-controls="about" role="tab" data-toggle="tab">About</a></li>
    </ul>

    <div class="tab-content">
      <div role="tabpanel" class="tab-pane fade in active" id="challenges">

        {{ if .Values.IsAdministrating }}
          <div style="position: relative;">
            <div style="position: absolute; top: 0; right: 0;">
              <a href="/cause/{{.Values.Cause.Id}}/challenges/new" class="btn btn-default">
                <span class="glyphicon glyphicon-plus text-primary"></span>
              </a>
            </div>
          </div>
        {{ end }}


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
          <div style="padding: 10px;">No challenges.</div>
        {{ end }}

      </div>
      <div role="tabpanel" class="tab-pane fade" id="about" style="padding: 10px;">
        <p>{{ .Values.Cause.Data.Description | format }}</p>
      </div>
    </div>

  </div>
  <div class="col-sm-4">
    <img src="{{ .Values.Cause.Info.IconURL }}"
        class="img-responsive img-rounded center-block"
        style="width:100%; margin-bottom:20px;"/>

    {{ if .Values.IsAdministrating }}
      <div class="list-group">
        <a class="list-group-item"
            href="/cause/{{ .Values.Cause.Id }}/admin/invite">Invite admin</a>
        <form method="post"
            onsubmit="return confirm('Are you sure you want to delete the cause?');">
          <input type="hidden" name="action" value="delete">
          <button type="submit" class="list-group-item">Delete</button>
        </form>
      </div>
    {{ end }}
  </div>
</div>

{{ template "footer" . }}`)
