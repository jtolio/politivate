package views

var _ = T.MustParse(`{{ template "header" (makemap "P" . "Title" "Set your district") }}

<h1>We need to figure out your location!</h1>

<p>We use this information solely to figure out what state and national
districts you're in.</p>

<br/>

<form class="form-horizontal" method="POST" name="location">
  <input type="hidden" id="inputLatitude" name="latitude">
  <input type="hidden" id="inputLongitude" name="longitude">

  <div class="col-sm-offset-2 col-sm-8">
    <div>
      <div class="input-group">
        <input type="text" class="form-control" id="inputAddress" name="address">
        <span class="input-group-btn">
          <button class="btn btn-primary" type="submit">Set</button>
        </span>
      </div>

      <div class="panel panel-default">
        <div id="placepicker" style="height: 300px;" class="panel-body"></div>
      </div>
    </div>
  </div>
</form>

<script>
function initPlacePicker() {
  $("#placepicker").locationpicker({
    location: {
      latitude: 0,
      longitude: 0
    },
    radius: 0,
    enableAutocomplete: true,
    addressFormat: "address",
    inputBinding: {
      latitudeInput: $('#inputLatitude'),
      longitudeInput: $('#inputLongitude'),
      locationNameInput: $('#inputAddress')
    }
  });
  navigator.geolocation.getCurrentPosition(function(location) {
    $("#placepicker").locationpicker("location", {
      latitude: location.coords.latitude,
      longitude: location.coords.longitude
    });
  });
}
</script>

{{.DeferredSources.Add "https://maps.google.com/maps/api/js?key=` + mapsAPIKey + `&libraries=places"}}
{{.DeferredSources.Add "/static/js/locationpicker.jquery.min.js"}}
{{.DeferredFuncs.Add "initPlacePicker"}}
{{ template "footer" . }}`)
