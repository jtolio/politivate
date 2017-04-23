_message:
	@echo "available targets: static compile routes serve deploy"

static:
	GOPATH=${PWD} go generate politivate.org/web/...

compile: static
	GOPATH=${PWD} go install -v politivate.org/web/...

routes: compile
	bin/local routes

serve: static _check_appengine
	GOPATH=${PWD} ${GCLOUD_SDK_ROOT}/bin/dev_appserver.py src/politivate.org/web/app/appengine

deploy: static
	GOPATH=${PWD} ${GCLOUD_SDK_ROOT}/bin/gcloud app deploy --project politiforce-150719 src/politivate.org/web/app/appengine/app.yaml

logs:
	GOPATH=${PWD} ${GCLOUD_SDK_ROOT}/bin/gcloud app logs tail --project politiforce-150719 -s default

_check_appengine:
	@if test -z "$$GCLOUD_SDK_ROOT"; then /bin/echo -e "\n\nError: please set GCLOUD_SDK_ROOT\n"; exit 1; fi

.PHONY: _message _check_appengine static compile routes serve deploy
