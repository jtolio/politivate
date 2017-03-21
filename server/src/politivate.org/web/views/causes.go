package views

var _ = T.MustParse(`{{ template "header" (makepair . "Causes") }}

<h1>Causes</h1>

{{ $ctx := .Ctx }}
{{ range $i, $c := .Values.Causes }}
  {{ if (ne $i 0) }}
    <div class="horizontal-line"></div>
  {{ end }}
  <a href="/cause/{{ $c.Id }}" class="large-button">
    <div class="media">
      <div class="media-left">
        <img class="media-object" style="border-radius: 10px;"
             src="{{ $c.Info.IconURL }}" alt="logo" width=64 height=64 />
      </div>
      <div class="media-body">
        <h4 class="media-heading">{{ $c.Info.Name }}</h4>
        {{ $count := ($c.UserCount $ctx) }}
        {{ if (eq $count 1) }}
          <p>1 follower</p>
        {{ else }}
          <p>{{$count}} followers</p>
        {{ end }}
      </div>
    </div>
  </a>
{{ end }}

{{ template "footer" . }}`)
