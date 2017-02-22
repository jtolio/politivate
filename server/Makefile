_message:
	@echo "available targets: static compile serve deploy rollback routes"

static:
	GOPATH=${PWD} go generate politivate.org/web/...

compile: static
	GOPATH=${PWD} go install -v politivate.org/web/...

routes: compile
	bin/local routes

serve: compile _check_appengine
	@echo
	@echo please see https://groups.google.com/d/msg/google-appengine-go/k-r_hJ7W1dw/jRkv8ZCMAgAJ
	@echo
	GOPATH=${PWD} ${GO_APPENGINE_ROOT}/goapp serve politivate.org/web/app/appengine

deploy: compile
	GOPATH=${PWD} ${GO_APPENGINE_ROOT}/goapp deploy politivate.org/web/app/appengine

rollback:
	${GO_APPENGINE_ROOT}/appcfg.py rollback src/politivate.org/web/app/appengine

_check_appengine:
	@if test -z "$$GO_APPENGINE_ROOT"; then /bin/echo -e "\n\nError: please set GO_APPENGINE_ROOT\n"; exit 1; fi

.PHONY: _message _check_appengine static compile serve deploy rollback
