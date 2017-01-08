package views

var _ = T.MustParse(`{{ template "header" (makepair . "New Challenge") }}

<p>User should enter:
<ul>
  <li>Title</li>
  <li>Description</li>
  <li>How many points</li>
</ul></p>

<p>User should pick:
<ul>
  <li>Phone call</li>
  <li>Location</li>
</ul></p>

<p><ul>
  <li>Enter a specific thing</li>
  <li>Use a database</li>
</ul></p>

<p><ul>
  <li>Deadline</li>
  <li>Time range to post</li>
</ul></p>

<p>Restrict to<ul>
  <li>States</li>
  <li>Districts</li>
  <li>Committees</li>
</ul></p>

{{ template "footer" . }}`)
