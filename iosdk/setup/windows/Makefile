IOSDK_VER?=master
WSK=https://github.com/apache/openwhisk-cli/releases/download/1.0.0/OpenWhisk_CLI-1.0.0-windows-amd64.zip
HERE:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: all
all: iosdk.exe wsk.exe
	docker run -v $(HERE):/mnt -e VER=$(IOSDK_VER) pagopa/builder-fpm-nsis bash /mnt/build.sh

iosdk.exe:
	GOOS=windows GOARCH=amd64 go build \
	-o iosdk.exe -ldflags "-X main.Version=$(IOSDK_VER)" ../.. 

wsk.exe:
	curl -sL $(WSK) >wsk.zip
	unzip wsk.zip wsk.exe

nsis.exe:
	curl -L $(NSIS) >nsis.exe

