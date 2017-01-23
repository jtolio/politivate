package static

//go:generate go install .../go-bindata...
//go:generate bash -c "cd ../../sass && pyscss application.scss > ../static/css/bootstrap.css"
//go:generate /usr/bin/env PATH=$GOPATH/bin go-bindata-assetfs -o bindata.go -pkg static --prefix ../../ ../../static/...

import (
	"net/http"

	"gopkg.in/webhelp.v1/whmux"
)

var (
	Handler http.Handler = whmux.RequireGet(http.FileServer(assetFS()))
)
