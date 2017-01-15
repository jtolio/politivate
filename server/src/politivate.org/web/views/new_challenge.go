package views

var _ = T.MustParse(`{{ template "header" (makepair . "New Challenge") }}

<h1>Create a new challenge</h1>

<form>
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
              <label class="btn btn-primary active">
                <input type="radio" name="type" id="type-phonecall" autocomplete="off" checked> Phone call
              </label>
              <label class="btn btn-primary">
                <input type="radio" name="type" id="type-location" autocomplete="off"> Location
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
      <div class="form-group">
        <label for="databaseInput">Phone number to call</label>
        <select class="form-control" name="database" id="databaseInput">
          <option>Call a specific number</option>
          <option>Call a local state representative</option>
        </select>
      </div>

      {{ template "input" (makemap "Field" "directphone" "Display" "Phone number" "Type" "tel" "Form" .Values.Form) }}
      {{ template "textarea" (makemap "Field" "directaddr" "Display" "Address" "Form" .Values.Form) }}

      <div class="form-group">
        <label>Add restriction</label>
        <div class="row">
          <div class="col-md-5">
            <select class="form-control">
              <option>State</option>
              <option>District</option>
              <option>Committee</option>
            </select>
          </div>
          <div class="col-md-5">
            <select class="form-control">
              <option>Alabama</option>
              <option>Alaska</option>
              <option>Arkansas</option>
            </select>
          </div>
          <div class="col-md-2">
            <button class="btn btn-default">Add</button>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="form-group">
    <button type="submit" class="btn btn-default">Submit</button>
  </div>
</form>

{{ template "footer" . }}`)
