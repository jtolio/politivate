_message:
	@echo "available targets: setup deploy run log flow"

setup:
	npm install

deploy:
	node_modules/.bin/react-native run-android

run:
	node_modules/.bin/react-native start

flow:
	node_modules/.bin/flow

log:
	node_modules/.bin/react-native log-android

.PHONY: _message setup deploy run log
