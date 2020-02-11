.PHONY: all
all: push-all
	echo TODO

.PHONY: build-ide-js-base
build-ide-js-base:
	docker build -t actionloop/ide-js-base ide-js-base

.PHONY: build-ide-js
build-ide-js: build-ide-js-base
	docker build -t actionloop/ide-js ide-js

.PHONY: push
push-all: build-ide-js
	docker push actionloop/ide-js


