VER=2020-11-16
REPO=https://github.com/sciabarracom/openwhisk
BRANCH=debugger
WHISK=pagopa/openwhisk-standalone:$(VER)
BUILDER=pagopa/builder-fpm-nsis:$(VER)
THEIA=pagopa/theia:$(VER)
PYTHONRUNTIME=pagopa/action-python-v3.7:$(VER)
NODEJSRUNTIME=pagopa/action-nodejs-v10:$(VER)

.PHONY: all builder openwhisk theia pythonruntime nodejsruntime

all: builder openwhisk theia pythonruntime nodejsruntime

openwhisk:
	test -d openwhisk || git clone $(REPO) openwhisk
	bash -c 'cd openwhisk ;\
	if test "$$(git rev-parse --abbrev-ref HEAD)" == $(BRANCH) ;\
	then git pull origin $(BRANCH) ;\
	else git checkout origin/$(BRANCH) -b $(BRANCH) ; fi'
	cd openwhisk ; ./gradlew :core:standalone:distDocker
	docker tag whisk/standalone $(WHISK)
	docker push $(WHISK)

builder:
	docker build -t $(BUILDER) builder
	docker push $(BUILDER)

theia:
	docker build -t $(THEIA) theia
	docker push $(THEIA)

pythonruntime:
	docker build -t $(PYTHONRUNTIME) pythonruntime
	docker push $(PYTHONRUNTIME)


nodejsruntime:
	docker build -t $(NODEJSRUNTIME) nodejsruntime
	docker push $(NODEJSRUNTIME)
