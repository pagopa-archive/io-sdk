AUTH=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP

.PHONY: deploy
deploy: setup deploy/send.zip
	wsk project deploy

.PHONY: undeploy
undeploy: setup
	wsk project undeploy

.PHONY:
setup:
	wsk property set --apihost http://localhost:3280 --auth $(AUTH)

deploy/send.zip: $(filter-out  %_test.go, $(wildcard src/send/*.go))
	zip - -r -j $^ | docker run -i openwhisk/action-golang-v1.11 -compile main >$@