.PHONY: all
all: push-all
	echo TODO

.PHONY: ide-js
ide-js:
	docker build -t actionloop/ide-js --squash ide-js

.PHONY: push
push-all: ide-js 
	docker push actionloop/ide-js

.PHONY: start-whisk
start-whisk:
	docker run -ti \
	-p 3232:3232 -p 3233:3233 \
	--rm --name openwhisk --hostname openwhisk \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-e CONTAINER_EXTRA_ENV=__OW_DEBUG_PORT=8081 \
	openwhisk/standalone:nightly

.PHONY: stop-whisk
stop-whisk:
	-docker exec openwhisk stop

.PHONY: start-ide
start-ide:
	docker run -ti $(OPENWHISK)\
	--rm --name ide-js \
	--add-host openwhisk:$(shell docker inspect openwhisk | jq -r '.[0].NetworkSettings.IPAddress') \
	-p 3000:3000 \
	-v /var/run/docker.sock:/var/run/docker.sock \
	--entrypoint=/bin/bash \
	actionloop/ide-js

.PHONY: stop-ide
stop-ide:
	-docker kill ide-js

.PHONY: stop
stop: stop-ide stop-whisk
