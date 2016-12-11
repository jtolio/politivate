package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jtolds/webhelp"
	"github.com/spacemonkeygo/flagfile"
	"github.com/spacemonkeygo/spacelog/setup"

	"politivate.org/web/app"
)

var (
	addr = flag.String("addr", ":8080", "address to listen on")
)

func main() {
	flagfile.Load()
	setup.MustSetup(os.Args[0])
	switch flag.Arg(0) {
	case "serve":
		panic(webhelp.ListenAndServe(*addr, app.RootHandler))
	case "routes":
		webhelp.PrintRoutes(os.Stdout, app.RootHandler)
	default:
		fmt.Printf("Usage: %s <serve|routes>\n", os.Args[0])
	}
}
