_message:
	@echo "available targets: static compile routes serve deploy"

static:
	cd web; go generate ./...

compile: static
	cd web; go build -v -o ../server ./app/local

routes: compile
	./server routes

serve: static _check_appengine
	${GCLOUD_SDK_ROOT}/bin/dev_appserver.py ./web

deploy: static
	${GCLOUD_SDK_ROOT}/bin/gcloud app deploy --project politiforce-150719 ./web/app.yaml

logs:
	${GCLOUD_SDK_ROOT}/bin/gcloud app logs tail --project politiforce-150719 -s default

_check_appengine:
	@if test -z "$$GCLOUD_SDK_ROOT"; then /bin/echo -e "\n\nError: please set GCLOUD_SDK_ROOT\n"; exit 1; fi

.PHONY: _message _check_appengine static compile routes serve deploy
