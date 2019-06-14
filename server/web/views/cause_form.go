package views

var _ = T.MustParse(`
{{ if .Error }}
  <div class="alert alert-danger" role="alert">{{ .Error }}</div>
{{ end }}

<form class="form-horizontal" method="POST">
  <div class="form-group">
    <label for="inputName" class="col-sm-2 control-label">Name</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="inputName" name="name"
             value="{{ (index .Form "name") }}">
    </div>
  </div>
  <div class="form-group">
    <label for="inputURL" class="col-sm-2 control-label">URL</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="inputURL" name="url"
             value="{{ (index .Form "url") }}">
    </div>
  </div>
  <div class="form-group">
    <label for="inputIconURL" class="col-sm-2 control-label">Icon URL</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="inputIconURL" name="icon_url"
             value="{{ (index .Form "icon_url") }}">
    </div>
  </div>
  <div class="form-group">
    <label for="inputShortDesc" class="col-sm-2 control-label">Short Summary</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" id="inputShortDesc" name="short_desc"
             value="{{ (index .Form "short_desc") }}">
    </div>
  </div>
  <div class="form-group">
    <label for="inputDescription" class="col-sm-2 control-label">Description</label>
    <div class="col-sm-10">
      <textarea class="form-control" id="inputDescription" rows="3"
                name="description"
            >{{ (index .Form "description") }}</textarea>
    </div>
  </div>
  <div class="form-group">
    <div class="col-sm-offset-2 col-sm-10">
      <button type="submit" class="btn btn-default">{{.Action}}</button>
    </div>
  </div>
</form>`)
