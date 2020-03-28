IMG=actionloop/ide-js:config

.PHONY: start
start:
	docker run -ti $(OPENWHISK)\
	--rm --name ide-js \
	--add-host openwhisk:$(shell docker inspect openwhisk | jq -r '.[0].NetworkSettings.IPAddress') \
	-p 3000:3000 \
	-v /var/run/docker.sock:/var/run/docker.sock \
	actionloop/ide-js

.PHONY: stop
stop:
	-docker kill ide-js

.PHONY: ide-js
ide-js:
	docker build -t $(IMG) ide-js

.PHONY: push
push: ide-js 
	docker push $(IMG)



