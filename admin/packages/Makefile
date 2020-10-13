.PHONY: deploy
deploy:
	wsk project deploy

.PHONY: undeploy
undeploy:
	wsk project undeploy

.PHONY: test
test: deploy
	if ! bats ../test ; then cat /tmp/debug.log ; exit 1 ; fi
