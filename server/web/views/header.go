package views

var _ = T.MustParse(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    {{ if .Title }}
      <title>{{ .Title }} - Politivate.org</title>
    {{ else }}
      <title>Politivate.org</title>
    {{ end }}
    <link href="https://fonts.googleapis.com/css?family=Questrial" rel="stylesheet">
    <link rel="stylesheet" href="/static/css/bootstrap.css">
    <!--[if lt IE 9]>
      <script src="//oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
      <script src="//oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
    <link rel="stylesheet" href="/static/css/site.css">
    <link rel="icon" type="image/png" href="/static/images/icon128.png"
          sizes="128x128" />
  </head>
  <body>
    <nav class="navbar navbar-default">
      <div class="container-fluid">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed"
              data-toggle="collapse" data-target="#navbar-collapse-region"
              aria-expanded="false">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="/">
            <img alt="Politivate.org" src="/static/images/header.svg"
              height="20">
          </a>
        </div>
        <div class="collapse navbar-collapse" id="navbar-collapse-region">
          <ul class="nav navbar-nav">
            {{ if .P.Beta }}
              <li{{if .Selected}}{{if (eq .Selected "causes")}} class="active"{{end}}{{end}}>
                <a href="/causes">Causes</a>
              </li>
            {{ end }}
            <li{{if .Selected}}{{if (eq .Selected "about")}} class="active"{{end}}{{end}}>
              <a href="/about">About</a>
            </li>
            {{ if .P.Beta }}
              <li{{if .Selected}}{{if (eq .Selected "get")}} class="active"{{end}}{{end}}>
                <a href="/get">Get the App</a>
              </li>
            {{ end }}
          </ul>
          <ul class="nav navbar-nav navbar-right">
            {{ if .P.User }}
              <li class="dropdown">
                <a href="#" class="dropdown-toggle" data-toggle="dropdown"
                    role="button" aria-haspopup="true"
                    aria-expanded="false">{{ .P.User.Name }}
                    <span class="caret"></span></a>
                <ul class="dropdown-menu">
                  {{ if .P.Beta }}
                    <li><a href="/profile">Profile</a></li>
                    {{ if .P.User.CanCreateCause }}
                      <li><a href="/causes/new">New Cause</a></li>
                    {{ end }}
                    <li role="separator" class="divider"></li>
                  {{ end }}
                  <li><a href="{{ .P.LogoutURL }}">Logout</a></li>
                </ul>
              </li>
            {{ else }}
              <li{{if .Selected}}{{if (eq .Selected "login")}} class="active"{{end}}{{end}}>
                <a href="{{ .P.LoginURL }}">Login</a>
              </li>
            {{ end }}
          </ul>
        </div>
      </div>
    </nav>
    <div class="container">`)