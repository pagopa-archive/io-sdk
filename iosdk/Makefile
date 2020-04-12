iosdk: ../.version $(shell ls *.go)
	go build -ldflags "-X main.Version=$(shell cat ../.version)" -o iosdk

.PHONY: test
test:
	go test -v
