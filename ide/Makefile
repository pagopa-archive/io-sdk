IOSDK_VER?=$(shell git rev-parse --abbrev-ref HEAD)
IMG=iosdk/iosdk-theia:$(IOSDK_VER)

.PHONY: push
push: iosdk-theia 
	docker push $(IMG)

.PHONY: iosdk-theia
iosdk-theia:
	docker build -t $(IMG) .

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
