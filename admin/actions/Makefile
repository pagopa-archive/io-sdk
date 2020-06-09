.PHONY: build
build: deploy/send.zip

.PHONY: deploy
deploy: build 
	wsk project deploy

.PHONY: undeploy
undeploy: setup
	wsk project undeploy

.PHONY: test
test: deploy
	# warm up 
	- wsk action invoke util/echo -r >/dev/null
	- bats test/cache.bats >/dev/null
	bats test

deploy/send.zip: $(filter-out  %_test.go, $(wildcard src/send/*.go))
	zip - -r -j $^ | docker run -i openwhisk/action-golang-v1.11 -compile main >$@

.PHONY: clean
clean:
	rm deploy/send.zip
