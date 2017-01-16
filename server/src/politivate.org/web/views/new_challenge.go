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
        <div class="row">
          <div class="col-md-4">
            <select class="form-control" id="restrictionType"
                    onchange="restrictionTypeChange(); return true;">
              <option value="state">State</option>
              <option value="district">District</option>
              <option value="housecommittee">House Committee</option>
              <option value="senatecommittee">Senate Committee</option>
            </select>
          </div>
          <div class="col-md-6">
            <select class="form-control" style="display: none;"
                    id="stateRestriction">
              <option>Alabama</option>
              <option>Alaska</option>
              <option>Arkansas</option>
            </select>
            <select class="form-control" style="display: none;"
                    id="districtRestriction">
              <option>UT-1</option>
              <option>UT-2</option>
              <option>UT-3</option>
            </select>
            <select class="form-control" style="display: none;"
                    id="houseCommitteeRestriction">
              <option>Agriculture</option>
              <option>Appropriations</option>
              <option>Budget</option>
            </select>
            <select class="form-control" style="display: none;"
                    id="senateCommitteeRestriction">
              <option>Aging</option>
              <option>Appropriations</option>
              <option>Budget</option>
            </select>
          </div>
          <div class="col-md-2">
            <button onclick="addRestriction(); return false;"
                    class="btn btn-default"
                    type="button">+</button>
          </div>
        </div>
        <div id="restrictionList"></div>
      </div>
    </div>

    {{ $challengeType := (or (index .Values.Form "type") "") }}
    <div class="col-md-4">
      <div class="form-group">
        <label for="typeInput">Challenge type</label><br/>
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
        <div class="form-group">
          <label for="locationDatabaseInput">Address to visit</label>
          <select class="form-control" id="locationDatabaseInput"
                  name="locationDatabase"
                  onchange="locationDatabaseChange(); return true;">
            <option value="us"
                {{if (eq $locationDatabase "us")}}selected{{end}}>
              Go to your local legislator's (US House or Senate) office</option>
            <option value="ushouse"
                {{if (eq $locationDatabase "ushouse")}}selected{{end}}>
              Go to your local legislator's (US House) office</option>
            <option value="ussenate"
                {{if (eq $locationDatabase "ussenate")}}selected{{end}}>
              Go to your local legislator's (US Senate) office</option>
            <option value="direct"
                {{if (eq $locationDatabase "direct")}}selected{{end}}>
              Go to a specific address</option>
          </select>
        </div>

        <div id="specificLocationSection" style="display: none;">
          {{ template "textarea" (makemap "Field" "directaddr" "Display" "Address" "Form" .Values.Form "Rows" 3) }}
        </div>
      </div>

      {{ template "optinput" (makemap "Field" "startdate" "Display" "Start date" "Type" "date" "Form" .Values.Form) }}
      {{ template "optinput" (makemap "Field" "deadline" "Display" "Deadline" "Type" "date" "Form" .Values.Form) }}
    </div>
  </div>

  <div class="form-group" style="text-align: right;">
    <button type="submit" class="btn btn-primary">Submit</button>
  </div>
</form>

{{ template "footerscripts" . }}
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
    "district": "#districtRestriction",
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

    for (var i = 0; i < restrictions.length; i++) {
      var p = $("<p>");
      if (i > 0) {
        p.append("OR ");
      }
      p.append(restrictions[i].type + ": " + restrictions[i].value + " (");
      p.append($("<a>")
          .attr("href", "#")
          .attr("onclick", "removeRestriction(" + i + "); return false;")
          .append("remove"));
      p.append(")");
      rl.append(p);
      rl.append($("<input>")
          .attr("type", "hidden")
          .attr("name", "restrictionType[" + i + "]")
          .attr("value", restrictions[i].type));
      rl.append($("<input>")
          .attr("type", "hidden")
          .attr("name", "restrictionValue[" + i + "]")
          .attr("value", restrictions[i].value));
    }
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

  $(challengeTypeChange);
  $(phoneDatabaseChange);
  $(locationDatabaseChange);
  $(restrictionTypeChange);
  $(updateRestrictions);
</script>
{{ template "footerdoc" . }}
`)
