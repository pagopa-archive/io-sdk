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

.PHONY: undeploy
undeploy:
	$(MAKE) -C packages undeploy
