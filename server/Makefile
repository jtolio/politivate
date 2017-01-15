_message:
	@echo "available targets: static compile run deploy"

static:
	GOPATH=${PWD} go generate politivate.org/web/...

compile: static
	GOPATH=${PWD} go install -v politivate.org/web/...

run: compile _check_appengine
	GOPATH=${PWD} ${GO_APPENGINE_ROOT}/goapp serve politivate.org/web/app/appengine

deploy: compile
	GOPATH=${PWD} ${GO_APPENGINE_ROOT}/goapp deploy politivate.org/web/app/appengine

_check_appengine:
	@if test -z "$$GO_APPENGINE_ROOT"; then /bin/echo -e "\n\nError: please set GO_APPENGINE_ROOT\n"; exit 1; fi

.PHONY: _message _check_appengine static compile run deploy
