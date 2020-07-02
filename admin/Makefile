IOSDK_VER?=$(shell git rev-parse --abbrev-ref HEAD)
IMG=pagopa/iosdk-openwhisk:$(IOSDK_VER)

.PHONE: build
build: public/build/bundle.js
	$(MAKE) -C actions
	docker build -t $(IMG) .

.PHONY: push
push:
	docker push $(IMG)

.PHONY: clean
clean:
	-$(MAKE) -C actions clean
	-docker rmi -f $(IMG)

.PHONY: deploy
deploy:
	$(MAKE) -C actions deploy

.PHONY: devel
devel:
	npm install
	npm run dev

.PHONY: start
start: ../iosdk/iosdk
	../iosdk/iosdk start --skip-ide --skip-pull-images	

.PHONY: stop
stop: ../iosdk/iosdk
	-../iosdk/iosdk stop

../iosdk/iosdk: build
	$(MAKE) -C ../iosdk

public/build/bundle.js: $(wildcard src/*.svelte)
	node -v | grep v12
	npm install
	npm run build


.PHONY: undeploy
undeploy:
	$(MAKE) -C actions undeploy
