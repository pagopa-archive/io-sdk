iosdk: $(shell ls *.go)
	go build -ldflags "-X main.Version=$(shell ../version.sh)" -o iosdk

.PHONY: test
test:
	go test -v | tee difftest.out
	python3 difftest.py difftest.out

