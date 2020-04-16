.PHONY: all linux mac windows
IOSDK_VER?=0.1-unknown

all: linux mac windows

linux:
	$(MAKE) IOSDK_VER=$(IOSDK_VER) -C linux

mac:
	if which pkgbuild ;\
	then $(MAKE) IOSDK_VER=$(IOSDK_VER) -C mac ;\
	else echo "Not a Mac, skipping" ; fi

windows:
	$(MAKE) IOSDK_VER=$(IOSDK_VER) -C windows

clean:
	-rm version linux/*.rpm linux/*.deb mac/*.pkg mac/*.zip windows/*.zip windows/*.exe
	-rm -Rvf linux/bin mac/bin
	-rm -Rvf mac/target mac/application

builder-fpm-nsis:
	docker build -t pagopa/builder-fpm-nsis .
	docker push pagopa/builder-fpm-nsis
