package static

//go:generate go install .../go-bindata...
//go:generate bash -c "cd ../../sass && pyscss application.scss > ../static/css/bootstrap-new.css"
//go:generate bash -c "cd ../../static/css && if ! [ -e bootstrap.css ] || ! diff -q bootstrap.css bootstrap-new.css; then cp bootstrap{-new,}.css; fi"
//go:generate rm -f ../../static/css/bootstrap-new.css
//go:generate /usr/bin/env PATH=$GOPATH/bin go-bindata-assetfs -o bindata.go -pkg static --prefix ../../ ../../static/...
//go:generate bash -c "if ! [ -e assets.go ] || ! diff -q bindata_assetfs.go assets.go; then cp bindata_assetfs.go assets.go; fi"
//go:generate rm -f bindata_assetfs.go

import (
	"net/http"

	"gopkg.in/webhelp.v1/whmux"
)

var (
	Handler http.Handler = whmux.RequireGet(http.FileServer(assetFS()))
)
