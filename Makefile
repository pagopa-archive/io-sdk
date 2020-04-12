.PHONY: all
all: .version io-sdk
	$(MAKE) -C backend
	$(MAKE) -C ides

.PHONY: .version
.version:
	@if test -z `git tag --points-at HEAD` ;\
	then git rev-parse --abbrev-ref HEAD >.version ;\
	else git tag --points-at HEAD | head -1 >.version ;\
	fi

io-sdk: .version main.go $(shell ls wskide/*.go)
	go build -ldflags "-X main.Version=$(shell cat  .version)" -o io-sdk

deploy:
	$(MAKE) -C backend deploy

.PHONY: test
test:
	cd wskide && go test -v

snapshot:
	git tag $(shell date +snap-%Y-%m%d-%H%M)
