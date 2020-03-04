IMG=actionloop/ide-js:latest

.PHONY: all
all: push-all
	echo TODO

.PHONY: ide-js
ide-js:
	docker build -t $(IMG) --squash ide-js

.PHONY: push
push-all: ide-js 
	docker push $(IMG)


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
