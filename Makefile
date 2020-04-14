.PHONY: branch build test release snapshot all
branch: 
	$(MAKE) build 
	#$(MAKE) test 
	$(MAKE) push

build:
	$(MAKE) -C admin
	$(MAKE) -C ide
	$(MAKE) -C iosdk

push:
	$(MAKE) -C admin push
	$(MAKE) -C ide push

test:
	bash test.sh

release:
	test -n "$(VER)"
	$(MAKE) IOSDK_VER=$(VER) build
	#$(MAKE) test
	$(MAKE) IOSDK_VER=$(VER) push
	$(MAKE) IOSDK_VER=$(VER) -C iosdk/setup

snapshot:
	make VER=$(shell date +%Y.%m%d.%H%M-snapshot) release
