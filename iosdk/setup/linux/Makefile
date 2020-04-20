WSK=https://github.com/apache/openwhisk-cli/releases/download/1.0.0/OpenWhisk_CLI-1.0.0-linux-amd64.tgz
IOSDK_VER?=0.0-unknown
HERE:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
.PHONY: all
all: bin/iosdk bin/wsk
	docker run --rm -v $(HERE):/mnt pagopa/builder-fpm-nsis bash /mnt/build.sh $(IOSDK_VER)

bin/iosdk:
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build \
	-o bin/iosdk -ldflags "-X main.Version=$(IOSDK_VER)" ../..

bin/wsk:
	mkdir -p bin
	curl -sL $(WSK) | tar xzvf - -C bin wsk

.PHONY: clean
clean:
	rm -rfv bin deb rpm

