WSK=https://github.com/apache/openwhisk-cli/releases/download/1.0.0/OpenWhisk_CLI-1.0.0-mac-amd64.zip
IOSDK_VER?=0.0-unknown

.PHONY: all prereq pkg

all: prereq pkg

prereq:
	# you need fpm installed with
	# sudo gem install fpm
	which fpm

pkg: bin/iosdk bin/wsk
	fpm -s dir -t osxpkg -n iosdk -v $(IOSDK_VER) ./bin=/usr

bin/iosdk:
	mkdir -p bin
	GOOS=darwin GOARCH=amd64 go build \
	-o bin/iosdk -ldflags "-X main.Version=$(IOSDK_VER)" ../.. 

bin/wsk:
	mkdir -p bin
	curl -sL $(WSK) >wsk.zip
	unzip wsk.zip wsk -d bin


