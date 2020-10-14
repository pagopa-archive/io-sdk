IOSDK_VER?=$(shell git rev-parse --abbrev-ref HEAD)
IMG=pagopa/iosdk-openwhisk:$(IOSDK_VER)

.PHONE: build
build: web/public/build/bundle.js
	docker build -t $(IMG) .

.PHONY: push
push:
	docker push $(IMG)

.PHONY: clean
clean:
	-docker rmi -f $(IMG)

.PHONY: deploy
deploy:
	$(MAKE) -C packages deploy

.PHONY: devel
devel:
	cd web && npm install && npm run dev

.PHONY: start
start: ../iosdk/iosdk
	../iosdk/iosdk debug wskprops
	../iosdk/iosdk start --skip-ide --skip-pull-images	

.PHONY: stop
stop: ../iosdk/iosdk
	-../iosdk/iosdk stop

../iosdk/iosdk: build
	$(MAKE) -C ../iosdk

web/public/build/bundle.js: $(wildcard web/src/*.svelte)
	node -v | grep v12
	cd web && npm install && npm run build

### Backend

.PHONY: deploy
deploy:
	wsk project deploy

.PHONY: undeploy
undeploy:
	wsk project undeploy

.PHONY: test
test: deploy
	if ! bats test ; then cat /tmp/debug.log ; exit 1 ; fi


.PHONY: nim_test
nim_test: nim_deploy
	wsk property set --auth $(shell nim auth current --auth) --apihost $(shell nim auth current --apihost)
	if ! bats test ; then echo "FAIL: check  /tmp/debug.log"; exit 1 ; fi

.PHONY: nim_deploy
nim_deploy:
	nim project deploy .
