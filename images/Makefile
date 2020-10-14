VER=2020.06.08
REPO=https://github.com/sciabarracom/openwhisk
BRANCH=debugger
WHISK=pagopa/openwhisk-standalone:$(VER)
BUILDER=pagopa/builder-fpm-nsis:$(VER)
THEIA=pagopa/theia:$(VER)

PYTHONRUNTIME=sciabarracom/actionloop-python-v3.7:2020-10-13
NODEJSRUNTIME=sciabarracom/action-nodejs-v10:2020-10-13

.PHONY: all builder openwhisk theia pythonruntime nodejsruntime

all: builder openwhisk theia pythonruntime nodejsruntime

openwhisk:
	test -d openwhisk || git clone $(REPO) openwhisk
	bash -c 'cd openwhisk ;\
	if test "$$(git rev-parse --abbrev-ref HEAD)" == $(BRANCH) ;\
	then git pull origin $(BRANCH) ;\
	else git checkout origin/$(BRANCH) -b $(BRANCH) ; fi'
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
