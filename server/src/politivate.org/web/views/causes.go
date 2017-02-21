package views

var _ = T.MustParse(`{{ template "header" (makepair . "Causes") }}

<h1>Causes</h1>

<ul>
{{ range $c := .Values.Causes }}
  <li><a href="/cause/{{ $c.Id }}"
      ><img width=24 height=24 src="{{ $c.Info.IconURL }}" /> {{ $c.Info.Name }}</a></li>
{{ end }}
</ul>

{{ template "footer" . }}`)
