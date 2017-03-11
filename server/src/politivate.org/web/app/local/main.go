package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spacemonkeygo/flagfile"
	"github.com/spacemonkeygo/spacelog/setup"
	"gopkg.in/webhelp.v1/whlog"
	"gopkg.in/webhelp.v1/whroute"

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
		panic(whlog.ListenAndServe(*addr, app.RootHandler))
	case "routes":
		whroute.PrintRoutes(os.Stdout, app.RootHandler)
	default:
		fmt.Printf("Usage: %s <serve|routes>\n", os.Args[0])
	}
}
