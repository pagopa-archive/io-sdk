.PHONY: deploy
deploy: deploy/send.zip
	wsk project deploy

.PHONY: undeploy
undeploy:
	wsk project undeploy

deploy/send.zip: $(filter-out  %_test.go, $(wildcard src/send/*.go))
	zip - -r -j $^ | docker run -i openwhisk/action-golang-v1.11 -compile main >$@