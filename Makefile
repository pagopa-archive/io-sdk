.PHONY: all
all: iosdk
	$(MAKE) -C wskide
	$(MAKE) -C backend
	$(MAKE) -C ides

iosdk: main.go $(shell ls wskide/*.go)
	go build -o iosdk

deploy:
	$(MAKE) -C backend deploy

.PHONY: test
test:
	cd wskide && go test -v
