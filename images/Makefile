VER=2020.04.17
REPO=https://github.com/sciabarracom/openwhisk
BRANCH=debugger
WHISK=pagopa/openwhisk-standalone:$(VER)
BUILDER=pagopa/builder-fpm-nsis:$(VER)
THEIA=pagopa/theia:$(VER)

.PHONY: all builder openwhisk theia

all: builder openwhisk theia

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
