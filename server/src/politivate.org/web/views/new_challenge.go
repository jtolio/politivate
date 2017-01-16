package views

var _ = T.MustParse(`{{ template "header" (makepair . "New Challenge") }}

<h1>Create a new challenge</h1>

<form name="newchallenge" method="POST">
  {{ define "input" }}
    <div class="form-group">
      <label for="{{.Field}}Input">{{.Display}}</label>
      <input type="{{.Type}}" class="form-control" id="{{.Field}}Input"
             name="{{.Field}}" placeholder="{{.Placeholder}}"
             value="{{index .Form .Field}}" />
    </div>
  {{ end }}

  {{ define "textarea" }}
    <div class="form-group">
      <label for="{{.Field}}Input">{{.Display}}</label>
      <textarea class="form-control" id="{{.Field}}Input" rows="3"
                name="{{.Field}}"
            >{{index .Form .Field}}</textarea>
    </div>
  {{ end }}

  <div class="row">
    <div class="col-md-6">

      {{ template "input" (makemap "Field" "title" "Display" "Title" "Type" "text" "Form" .Values.Form) }}
      {{ template "textarea" (makemap "Field" "description" "Display" "Description" "Form" .Values.Form) }}
      {{ template "input" (makemap "Field" "points" "Display" "Points" "Type" "number" "Placeholder" "10" "Form" .Values.Form) }}

      <div class="row">
        <div class="col-md-4">
          <div class="form-group">
            <label for="typeInput">Challenge type</label><br/>
            <div class="btn-group" data-toggle="buttons">
              <label class="btn btn-default active">
                <input type="radio" name="type" id="type-phonecall"
                       autocomplete="off" value="phonecall" checked
                       onchange="challengeTypeChange(); return true;">
                  Phone call
              </label>
              <label class="btn btn-default">
                <input type="radio" name="type" id="type-location"
                       autocomplete="off" value="location"
                       onchange="challengeTypeChange(); return true;">
                  Location
              </label>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          {{ template "input" (makemap "Field" "deadline" "Display" "Deadline" "Type" "date" "Form" .Values.Form) }}
        </div>
        <div class="col-md-4">
          {{ template "input" (makemap "Field" "startdate" "Display" "Start date" "Type" "date" "Form" .Values.Form) }}
        </div>
      </div>
    </div>

    <div class="col-md-6">

      <div id="phoneDatabaseSection" style="display: none;">
        <div class="form-group">
          <label for="phoneDatabaseInput">Phone number to call</label>
          <select class="form-control" id="phoneDatabaseInput"
                  name="phoneDatabase"
                  onchange="phoneDatabaseChange(); return true;">
            <option value="national">Call a local state representative</option>
            <option value="specific">Call a specific number</option>
          </select>
        </div>

        <div id="specificPhoneSection" style="display: none;">
          {{ template "input" (makemap "Field" "directphone" "Display" "Phone number" "Type" "tel" "Form" .Values.Form) }}
        </div>
      </div>

      <div id="locationDatabaseSection" style="display: none;">
        <div class="form-group">
          <label for="locationDatabaseInput">Address to visit</label>
          <select class="form-control" id="locationDatabaseInput"
                  name="locationDatabase"
                  onchange="locationDatabaseChange(); return true;">
            <option value="national">Go to a local state representative's
                office</option>
            <option value="specific">Go to a specific address</option>
          </select>
        </div>

        <div id="specificLocationSection" style="display: none;">
          {{ template "textarea" (makemap "Field" "directaddr" "Display" "Address" "Form" .Values.Form) }}
        </div>
      </div>

      <div class="form-group">
        <label>Restrictions</label>
        <div id="restrictionList"></div>
        <div class="row">
          <div class="col-md-3">
            <select class="form-control" id="restrictionType"
                    onchange="restrictionTypeChange(); return true;">
              <option value="state">State</option>
              <option value="district">District</option>
              <option value="committee">Committee</option>
            </select>
          </div>
          <div class="col-md-5">
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
                    id="committeeRestriction">
              <option>Ways and Means</option>
              <option>Science</option>
              <option>Energy</option>
            </select>
          </div>
          <div class="col-md-4">
            <button onclick="addRestriction(); return false;"
                    class="btn btn-default"
                    type="button">Add Restriction</button>
          </div>
        </div>
      </div>
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
      "specific": "#specificPhoneSection"});
  }

  function locationDatabaseChange() {
    fieldChange($("#locationDatabaseInput").val(), {
      "specific": "#specificLocationSection"});
  }

  var restrictionListMap = {
    "state": "#stateRestriction",
    "district": "#districtRestriction",
    "committee": "#committeeRestriction"};

  function restrictionTypeChange() {
    fieldChange($("#restrictionType").val(), restrictionListMap);
  }

  var restrictions = [];

  function updateRestrictions() {
    var rl = $("#restrictionList");
    if (restrictions.length == 0) {
      rl.html("<p>No restrictions!</p>");
      return;
    }
    rl.empty();
    for (var i = 0; i < restrictions.length; i++) {
      var p = $("<p>");
      if (i > 0) {
        p.append("OR ");
      }
      p.append(restrictions[i].type + ": " + restrictions[i].value + " (");
      var a = $("<a>");
      a.attr("href", "#")
       .attr("onclick", "removeRestriction(" + i + "); return false;")
       .append("remove");
      p.append(a);
      p.append(")");
      rl.append(p);
      var inp = $("<input>");
      inp.attr("type", "hidden")
         .attr("name", "restriction[]")
         .attr("value", restrictions[i].type + ":" + restrictions[i].value);
      rl.append(inp);
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
