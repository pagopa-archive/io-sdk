.PHONY: all linux mac windows
IOSDK_VER?=0.1-unknown

all: linux mac windows

linux:
	$(MAKE) IOSDK_VER=$(IOSDK_VER) -C linux

mac:
	$(MAKE) IOSDK_VER=$(IOSDK_VER) -C mac

windows:
	$(MAKE) IOSDK_VER=$(IOSDK_VER) -C windows

clean:
	-rm version linux/*.rpm linux/*.deb mac/*.pkg mac/*.zip windows/*.zip windows/*.exe
	-rm -Rvf linux/bin mac/bin

builder-fpm-nsis:
	docker build -t iosdk/builder-fpm-nsis .
	docker push iosdk/builder-fpm-nsis
