function doAction(action, message) {
  if (message && !confirm(message)) {
    return;
  }
  var form = $("<form>");
  form.attr("method", "post");
  var input = $("<input>");
  input.attr("type", "hidden");
  input.attr("name", "action");
  input.attr("value", action);
  form.append(input);
  $(document.body).append(form);
  form.submit();
}
