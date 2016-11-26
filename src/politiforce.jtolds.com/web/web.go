package web

import (
	"net/http"

	"github.com/jtolds/webhelp"
	"golang.org/x/net/context"
)

func challenges(ctx context.Context, w webhelp.ResponseWriter, r *http.Request) error {
	w.Write([]byte("hello 2"))
	return nil
}

func init() {
	http.Handle("/", webhelp.Base{Root: webhelp.DirMux{
		"challenges": webhelp.HandlerFunc(challenges),
	}})
}
