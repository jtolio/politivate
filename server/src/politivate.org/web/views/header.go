package views

var _ = T.MustParse(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    {{ if .Second }}
      <title>{{ .Second }} - Politivate.org</title>
    {{ else }}
      <title>Politivate.org</title>
    {{ end }}
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
              height="40">
          </a>
        </div>
        <div class="collapse navbar-collapse" id="navbar-collapse-region">
          <ul class="nav navbar-nav">
            <li><a href="/causes">Causes</a></li>
            <li><a href="/about">About</a></li>
            <li><a href="/get">Get the App</a></li>
          </ul>
          <ul class="nav navbar-nav navbar-right">
            {{ if .First.User }}
              <li class="dropdown">
                <a href="#" class="dropdown-toggle" data-toggle="dropdown"
                    role="button" aria-haspopup="true"
                    aria-expanded="false">{{ .First.User.Name }} <span class="caret"></span></a>
                <ul class="dropdown-menu">
                  <li><a href="/profile">Profile</a></li>
                  <li role="separator" class="divider"></li>
                  <li><a href="{{ .First.LogoutURL }}">Logout</a></li>
                </ul>
              </li>
            {{ else }}
              <li><a href="{{ .First.LoginURL }}">Login</a></li>
            {{ end }}
          </ul>
        </div>
      </div>
    </nav>
    <div class="container">`)
