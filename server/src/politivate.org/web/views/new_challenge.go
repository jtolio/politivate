package views

var _ = T.MustParse(`{{ template "header" (makepair . "New Challenge") }}

<h1>Create a new challenge</h1>

{{ if .Values.Error }}
  <div class="alert alert-danger" role="alert">{{ .Values.Error }}</div>
{{ end }}

{{ define "input" }}
  <div class="form-group">
    <label for="{{.Field}}Input">{{.Display}}</label>
    <input type="{{.Type}}" class="form-control" id="{{.Field}}Input"
           name="{{.Field}}" placeholder="{{.Placeholder}}"
           value="{{index .Form .Field}}" />
  </div>
{{ end }}

{{ define "optinput" }}
  <div class="form-group">
    <label for="{{.Field}}Input">{{.Display}}</label>
    <div class="input-group">
      <span class="input-group-addon">
        <input type="checkbox" id="{{.Field}}OptInput"
               name="{{.Field}}Enabled"
               {{if (index .Form (printf "%s%s" .Field "Enabled"))}}checked{{end}} />
      </span>
      <input type="{{.Type}}" class="form-control" id="{{.Field}}Input"
             name="{{.Field}}" placeholder="{{.Placeholder}}"
             value="{{index .Form .Field}}" />
    </div>
  </div>
{{ end }}

{{ define "textarea" }}
  <div class="form-group">
    <label for="{{.Field}}Input">{{.Display}}</label>
    <textarea class="form-control" id="{{.Field}}Input" rows="{{.Rows}}"
              name="{{.Field}}"
          >{{index .Form .Field}}</textarea>
  </div>
{{ end }}

<form name="newchallenge" method="POST">
  <div class="row">
    <div class="col-md-8">

      {{ template "input" (makemap "Field" "title" "Display" "Title" "Type" "text" "Form" .Values.Form) }}
      {{ template "input" (makemap "Field" "points" "Display" "Points" "Type" "number" "Placeholder" "10" "Form" .Values.Form) }}
      {{ template "textarea" (makemap "Field" "description" "Display" "Description" "Form" .Values.Form "Rows" 10) }}

      <div class="form-group">
        <label>Restrictions</label>
        <div id="restrictionList"></div>
        <div class="row">
          <div class="col-md-4">
            <select class="form-control" id="restrictionType"
                    onchange="restrictionTypeChange(); return true;">
              <option value="state">State</option>
              <option value="housecommittee">House Committee</option>
              <option value="senatecommittee">Senate Committee</option>
            </select>
          </div>
          <div class="col-md-6">
            <select class="form-control" style="display: none;"
                    id="stateRestriction">
              <option>Alabama</option>
              <option>Alaska</option>
              <option>American Samoa</option>
              <option>Arizona</option>
              <option>Arkansas</option>
              <option>California</option>
              <option>Colorado</option>
              <option>Connecticut</option>
              <option>Delaware</option>
              <option>District of Columbia</option>
              <option>Florida</option>
              <option>Georgia</option>
              <option>Guam</option>
              <option>Hawaii</option>
              <option>Idaho</option>
              <option>Illinois</option>
              <option>Indiana</option>
              <option>Iowa</option>
              <option>Kansas</option>
              <option>Kentucky</option>
              <option>Louisiana</option>
              <option>Maine</option>
              <option>Maryland</option>
              <option>Massachusetts</option>
              <option>Michigan</option>
              <option>Minnesota</option>
              <option>Mississippi</option>
              <option>Missouri</option>
              <option>Montana</option>
              <option>Nebraska</option>
              <option>Nevada</option>
              <option>New Hampshire</option>
              <option>New Jersey</option>
              <option>New Mexico</option>
              <option>New York</option>
              <option>North Carolina</option>
              <option>North Dakota</option>
              <option>Northern Mariana Islands</option>
              <option>Ohio</option>
              <option>Oklahoma</option>
              <option>Oregon</option>
              <option>Pennsylvania</option>
              <option>Philippines</option>
              <option>Puerto Rico</option>
              <option>Rhode Island</option>
              <option>South Carolina</option>
              <option>South Dakota</option>
              <option>Tennessee</option>
              <option>Texas</option>
              <option>U.S. Virgin Islands</option>
              <option>Utah</option>
              <option>Vermont</option>
              <option>Virginia</option>
              <option>Washington</option>
              <option>West Virginia</option>
              <option>Wisconsin</option>
              <option>Wyoming</option>
            </select>
            <select class="form-control" style="display: none;"
                    id="houseCommitteeRestriction">
                            <option>Agriculture</option>
              <option>Appropriations</option>
              <option>Armed Services</option>
              <option>Budget</option>
              <option>Economic</option>
              <option>Education and the Workforce</option>
              <option>Energy and Commerce</option>
              <option>Ethics</option>
              <option>Financial Services</option>
              <option>Foreign Affairs</option>
              <option>Homeland Security</option>
              <option>House Administration</option>
              <option>Intelligence</option>
              <option>Judiciary</option>
              <option>Library</option>
              <option>Natural Resources</option>
              <option>Oversight and Government Reform</option>
              <option>Printing</option>
              <option>Rules</option>
              <option>Science, Space, and Technology</option>
              <option>Small Business</option>
              <option>Taxation</option>
              <option>Transportation and Infrastructure</option>
              <option>Veterans' Affairs</option>
              <option>Ways and Means</option>
            </select>
            <select class="form-control" style="display: none;"
                    id="senateCommitteeRestriction">
              <option>Aging</option>
              <option>Agriculture, Nutrition, and Forestry</option>
              <option>Appropriations</option>
              <option>Armed Services</option>
              <option>Banking, Housing, and Urban Affairs</option>
              <option>Budget</option>
              <option>Commerce, Science, and Transportation</option>
              <option>Economic</option>
              <option>Energy and Natural Resources</option>
              <option>Environment and Public Works</option>
              <option>Ethics</option>
              <option>Finance</option>
              <option>Foreign Relations</option>
              <option>Health, Education, Labor, and Pensions</option>
              <option>Homeland Security and Governmental Affairs</option>
              <option>Indian Affairs</option>
              <option>Intelligence</option>
              <option>Judiciary</option>
              <option>Library</option>
              <option>Printing</option>
              <option>Rules and Administration</option>
              <option>Small Business and Entrepreneurship</option>
              <option>Taxation</option>
              <option>Veterans' Affairs</option>
            </select>
          </div>
          <div class="col-md-2">
            <button onclick="addRestriction(); return false;"
                    class="btn btn-default"
                    type="button">+</button>
          </div>
        </div>
      </div>
    </div>

    <div class="col-md-4">
      {{ $challengeType := (or (index .Values.Form "type") "") }}
      <div class="form-group">
        <label>Challenge type</label><br/>
        <div class="btn-group" data-toggle="buttons">
          <label class="btn btn-default{{if (or (not $challengeType) (eq $challengeType "phonecall"))}} active{{end}}">
            <input type="radio" name="type" id="type-phonecall"
                   autocomplete="off" value="phonecall"
                   {{if (or (not $challengeType) (eq $challengeType "phonecall"))}}checked{{end}}
                   onchange="challengeTypeChange(); return true;">
              Phone call
          </label>
          <label class="btn btn-default{{if (eq $challengeType "location")}} active{{end}}">
            <input type="radio" name="type" id="type-location"
                   autocomplete="off" value="location"
                   {{if (eq $challengeType "location")}}checked{{end}}
                   onchange="challengeTypeChange(); return true;">
              Location
          </label>
        </div>
      </div>

      {{ $phoneDatabase := (or (index .Values.Form "phoneDatabase") "") }}
      <div id="phoneDatabaseSection" style="display: none;">
        <div class="form-group">
          <label for="phoneDatabaseInput">Phone number to call</label>
          <select class="form-control" id="phoneDatabaseInput"
                  name="phoneDatabase"
                  onchange="phoneDatabaseChange(); return true;">
            <option value="us"
                {{if (eq $phoneDatabase "us")}}selected{{end}}>
              Call your local legislator in the US House or Senate</option>
            <option value="ushouse"
                {{if (eq $phoneDatabase "ushouse")}}selected{{end}}>
              Call your local legislator in the US House</option>
            <option value="ussenate"
                {{if (eq $phoneDatabase "ussenate")}}selected{{end}}>
              Call your local legislator in the US Senate</option>
            <option value="direct"
                {{if (eq $phoneDatabase "direct")}}selected{{end}}>
              Call a specific number</option>
          </select>
        </div>

        <div id="specificPhoneSection" style="display: none;">
          {{ template "input" (makemap "Field" "directphone" "Display" "Phone number" "Type" "tel" "Form" .Values.Form) }}
        </div>
      </div>

      {{ $locationDatabase := (or (index .Values.Form "locationDatabase") "") }}
      <div id="locationDatabaseSection" style="display: none;">
        <div class="form-group" style="display: none;">
          <label for="locationDatabaseInput">Address to visit</label>
          <select class="form-control" id="locationDatabaseInput"
                  name="locationDatabase"
                  onchange="locationDatabaseChange(); return true;">
            <option value="direct"
                {{if (eq $locationDatabase "direct")}}selected{{end}}>
              Go to a specific address</option>
            <option value="us"
                {{if (eq $locationDatabase "us")}}selected{{end}}>
              Go to your local legislator's (US House or Senate) office</option>
            <option value="ushouse"
                {{if (eq $locationDatabase "ushouse")}}selected{{end}}>
              Go to your local legislator's (US House) office</option>
            <option value="ussenate"
                {{if (eq $locationDatabase "ussenate")}}selected{{end}}>
              Go to your local legislator's (US Senate) office</option>
          </select>
        </div>

        <div id="specificLocationSection" style="display: none;">
          <div class="form-group">
            <label for="directaddrInput">Address</label>
            <div style="width: 300px;">
              <input type="text" class="form-control" id="directaddrInput"
                     name="directaddr" value="{{index .Values.Form "directaddr"}}">
              <input type="hidden" id="directlatInput" name="directlat"
                     value="{{index .Values.Form "directlat"}}">
              <input type="hidden" id="directlonInput" name="directlon"
                     value="{{index .Values.Form "directlon"}}">

              <div class="panel panel-default">
                <div id="directaddrPicker" style="height: 200px;" class="panel-body"></div>
              </div>
              <input type="number" class="form-control" id="directradiusInput"
                     name="directradius" step="20"
                     value="{{index .Values.Form "directradius"}}">
            </div>
          </div>
        </div>
      </div>

      {{ $dateType := (or (index .Values.Form "dateType") "") }}
      <div class="form-group">
        <label>Date type</label><br/>
        <div class="btn-group" data-toggle="buttons">
          <label class="btn btn-default{{if (or (not $dateType) (eq $dateType "none"))}} active{{end}}">
            <input type="radio" name="dateType" id="dateType-none"
                   autocomplete="off" value="none"
                   {{if (or (not $dateType) (eq $dateType "none"))}}checked{{end}}
                   onchange="dateTypeChange(); return true;">
              None
          </label>
          <label class="btn btn-default{{if (eq $dateType "deadline")}} active{{end}}">
            <input type="radio" name="dateType" id="dateType-deadline"
                   autocomplete="off" value="deadline"
                   {{if (eq $dateType "deadline")}}checked{{end}}
                   onchange="dateTypeChange(); return true;">
              Deadline
          </label>
          <label class="btn btn-default{{if (eq $dateType "event")}} active{{end}}">
            <input type="radio" name="dateType" id="dateType-event"
                   autocomplete="off" value="event"
                   {{if (eq $dateType "event")}}checked{{end}}
                   onchange="dateTypeChange(); return true;">
              Event
          </label>
        </div>
      </div>

      <div id="eventStartSection" style="display: none;">
        {{ template "input" (makemap "Field" "eventStart" "Display" "Start time" "Type" "datetime-local" "Form" .Values.Form) }}
      </div>
      <div id="eventEndSection" style="display: none;">
        {{ template "input" (makemap "Field" "eventEnd" "Display" (safehtml "<div id=\"eventEndLabel\">End time</div>") "Type" "datetime-local" "Form" .Values.Form) }}
      </div>
    </div>
  </div>

  <div class="form-group" style="text-align: right;">
    <button type="submit" class="btn btn-primary">Submit</button>
  </div>
</form>

{{ template "footerscripts" . }}
<script type="text/javascript" src="https://maps.google.com/maps/api/js?key=` + mapsAPIKey + `&libraries=places"></script>
<script src="/static/js/locationpicker.jquery.min.js"></script>
<script>
  function fieldChange(val, opts) {
    $.each(opts, function(optval, optselect) {
      if (optval == val) {
        $(optselect).show();
      } else {
        $(optselect).hide();
      }
    });
  }

  function challengeTypeChange() {
    fieldChange(document.forms["newchallenge"]["type"].value, {
      "phonecall": "#phoneDatabaseSection",
      "location": "#locationDatabaseSection"});
    if (document.forms["newchallenge"]["type"].value == "location") {
      setupAddress();
    }
  }

  function dateTypeChange() {
    switch (document.forms["newchallenge"]["dateType"].value) {
      case "deadline":
        $("#eventStartSection").hide();
        $("#eventEndSection").show();
        $("#eventEndLabel").text("Deadline");
        showEnd = true;
        break;
      case "event":
        $("#eventStartSection").show();
        $("#eventEndSection").show();
        $("#eventEndLabel").text("End time");
        break;
      default:
        $("#eventStartSection").hide();
        $("#eventEndSection").hide();
        break;
    }
  }

  function phoneDatabaseChange() {
    fieldChange($("#phoneDatabaseInput").val(), {
      "direct": "#specificPhoneSection"});
  }

  function locationDatabaseChange() {
    fieldChange($("#locationDatabaseInput").val(), {
      "direct": "#specificLocationSection"});
  }

  var restrictionListMap = {
    "state": "#stateRestriction",
    "housecommittee": "#houseCommitteeRestriction",
    "senatecommittee": "#senateCommitteeRestriction"};

  function restrictionTypeChange() {
    fieldChange($("#restrictionType").val(), restrictionListMap);
  }

  var restrictions = [
    {{ range $i, $e := .Values.Restrictions }}
    {{if $i}},{{end}}{"type": "{{$e.Type}}", "value": "{{$e.Value}}"}
    {{ end }}
  ];

  var restrictionTypeName = {
    "housecommittee": "House Committee",
    "senatecommittee": "Senate Committee",
    "state": "State"
  }

  function updateRestrictions() {
    var rl = $("#restrictionList");
    rl.empty();
    rl.append($("<input>")
        .attr("type", "hidden")
        .attr("name", "restrictionLength")
        .attr("value", "" + restrictions.length));

    if (restrictions.length == 0) {
      rl.append("<p>No restrictions!</p>");
      return;
    }

    var list = $("<ul>");
    for (var i = 0; i < restrictions.length; i++) {
      var item = $("<li>");
      if (i > 0) {
        item.append("OR ");
      }
      item.append(restrictionTypeName[restrictions[i].type] + ": " +
                  restrictions[i].value + " (");
      item.append($("<a>")
          .attr("href", "#")
          .attr("onclick", "removeRestriction(" + i + "); return false;")
          .append("remove"));
      item.append(")");
      item.append($("<input>")
          .attr("type", "hidden")
          .attr("name", "restrictionType[" + i + "]")
          .attr("value", restrictions[i].type));
      item.append($("<input>")
          .attr("type", "hidden")
          .attr("name", "restrictionValue[" + i + "]")
          .attr("value", restrictions[i].value));
      list.append(item);
    }
    rl.append(list);
  }

  function parseReal(val) {
    if (val.length == 0) {
      throw "no digits";
    }
    return parseFloat(val);
  }

  function addRestriction() {
    var rt = $("#restrictionType").val();
    var val = $(restrictionListMap[rt]).val();
    for (var i = 0; i < restrictions.length; i++) {
      if (restrictions[i].type == rt && restrictions[i].value == val) {
        return;
      }
    }
    restrictions.push({"type": rt, "value": val});
    updateRestrictions();
  }

  function removeRestriction(idx) {
    restrictions.splice(idx, 1);
    updateRestrictions();
  }

  var addressSetUp = false;
  function setupAddress() {
    if (addressSetUp) { return; }
    addressSetUp = true;

    var options = {
      enableAutocomplete: true,
      addressFormat: "address",
      inputBinding: {
        latitudeInput: $("#directlatInput"),
        longitudeInput: $("#directlonInput"),
        locationNameInput: $("#directaddrInput"),
        radiusInput: $("#directradiusInput")
      }
    };

    try {
      options.location = {
        latitude: parseReal($("#directlatInput").val()),
        longitude: parseReal($("#directlonInput").val())
      };
    } catch(err) {
      options.location = {
        latitude: 0,
        longitude: 0
      };
    }

    try {
      options.radius = parseReal($("#directradiusInput").val());
    } catch(err) {
      options.radius = 200;
    }

    $("#directaddrPicker").locationpicker(options);
  }

  $(challengeTypeChange);
  $(dateTypeChange);
  $(phoneDatabaseChange);
  $(locationDatabaseChange);
  $(restrictionTypeChange);
  $(updateRestrictions);
</script>
{{ template "footerdoc" . }}
`)
