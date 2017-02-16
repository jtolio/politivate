package views

var _ = T.MustParse(`{{ template "header" (makepair . "Set your district") }}

<h1>We need to figure out your location!</h1>

<p>We use this information solely to figure out what state and national
districts you're in.</p>

<form class="form-horizontal" method="POST" name="location">
  <div class="form-group">
    <label for="inputLatitude" class="col-sm-2 control-label">Latitude</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="inputLatitude"
             name="latitude">
    </div>
  </div>
  <div class="form-group">
    <label for="inputLongitude" class="col-sm-2 control-label">Longitude</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="inputLongitude"
             name="longitude">
    </div>
  </div>
  <div class="form-group">
    <div class="col-sm-offset-2 col-sm-10">
      <button type="submit" class="btn btn-default">Set</button>
    </div>
  </div>
</form>

{{ template "footerscripts" . }}
<script>
$(function() {
  navigator.geolocation.getCurrentPosition(function(location) {
    document.forms["location"].elements["longitude"].value = location.coords.longitude;
    document.forms["location"].elements["latitude"].value = location.coords.latitude;
  });
})
</script>
{{ template "footerdoc" . }}`)
