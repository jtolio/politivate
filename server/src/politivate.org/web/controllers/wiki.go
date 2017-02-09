package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/webhelp.v1/whauth"
	"gopkg.in/webhelp.v1/whredir"

	"politivate.org/web/secrets"
)

func init() {
	mux["wiki"] = whauth.RequireBasicAuth(whredir.RequireNextSlash(
		http.HandlerFunc(wikiHandler)), "Politivate Wiki", secrets.WikiAuth)
}

func wikiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><a href="https://app.nuclino.com/teams/13:26915">
  The wiki has moved to Nuclino</a>. Please ask JT
  (<a href="mailto:hello@jtolds.com">hello@jtolds.com</a>) for an invite.
  <html>`)
}
