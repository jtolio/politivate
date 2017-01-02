package views

var _ = T.MustParse(`{{ template "header" (makepair . "") }}
<style>
  .flag-jumbotron {
    background: url(/static/images/flag.jpg) no-repeat center;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
    padding-top: 400px;
    text-align: right;
  }
</style>
<div class="jumbotron flag-jumbotron">
  <p><a class="btn btn-primary btn-lg" href="/about" role="button">Make a difference</a></p>
</div>

{{ template "footer" . }}`)
