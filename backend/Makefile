IMG=actionloop/iosdk:config

.PHONY: deploy
deploy:
	$(MAKE) -C actions deploy

.PHONY: devel
devel:
	npm install
	npm run dev

.PHONY: start
start:
	docker run -ti \
	-p 3280:3280 -p 3281:3281 \
	--rm --name openwhisk --hostname openwhisk \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-e CONTAINER_EXTRA_ENV=__OW_DEBUG_PORT=8081 \
	$(IMG)

.PHONY: stop
stop:
	-docker exec openwhisk stop

.PHONE: build
build:
	node -v | grep v10
	npm install
	npm run build
	docker build -t $(IMG) .

push: build
	docker push $(IMG)

.PHONY: undeploy
undeploy:
	$(MAKE) -C actions undeploy
