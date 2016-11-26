package web

import (
	"net/http"

	"github.com/jtolds/webhelp"
)

func init() {
	http.HandleFunc("/", webhelp.DirMux{})
}
