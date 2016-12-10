package main

import (
	"flag"
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
	panic(webhelp.ListenAndServe(*addr, app.RootHandler))
}
