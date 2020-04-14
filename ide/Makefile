IOSDK_VER?=$(shell git rev-parse --abbrev-ref HEAD)
IMG=iosdk/iosdk-theia:$(IOSDK_VER)

.PHONY: build
build:
	docker build -t $(IMG) .

.PHONY: push
push: build 
	docker push $(IMG)

.PHONY: start
start:
	docker run -ti \
	--rm --name iosdk-theia \
	--add-host openwhisk:$(shell docker inspect iosdk-openwhisk | jq -r '.[0].NetworkSettings.IPAddress') \
	-p 3000:3000 \
	-v /var/run/docker.sock:/var/run/docker.sock \
	$(IMG)

.PHONY: stop
stop:
	-docker kill iosdk-theia