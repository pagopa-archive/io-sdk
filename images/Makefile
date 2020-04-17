VER=2020.04.17
WHISK=pagopa/openwhisk-standalone:$(VER)
BUILDER=pagopa/builder-fpm-nsis:$(VER)
THEIA=pagopa/theia:$(VER)

.PHONY: all builder openwhisk theia

all: builder openwhisk theia

openwhisk:
	docker build -t $(WHISK) openwhisk
	docker push $(WHISK)

builder:
	docker build -t $(BUILDER) builder
	docker push $(BUILDER)

theia:
	docker build -t $(THEIA) theia
	docker push $(THEIA)
