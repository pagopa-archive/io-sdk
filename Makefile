VER?=$(shell git tag --points-at HEAD | head -1)

.PHONY: branch build test release snapshot all
branch: 
	$(MAKE) build 
	$(MAKE) test 
	$(MAKE) push

release:
	test -n "$(VER)"
	$(MAKE) IOSDK_VER=$(VER) build
	$(MAKE) test
	$(MAKE) IOSDK_VER=$(VER) push
	$(MAKE) IOSDK_VER=$(VER) -C iosdk/setup/linux
	$(MAKE) IOSDK_VER=$(VER) -C iosdk/setup/windows

release_mac:
	test -n "$(VER)"
	$(MAKE) IOSDK_VER=$(VER) -C iosdk/setup/mac


clean:
	$(MAKE) -C admin clean
	$(MAKE) -C ide clean
	$(MAKE) -C iosdk clean
	
build:
	$(MAKE) -C admin
	$(MAKE) -C ide
	$(MAKE) -C iosdk

push:
	docker login
	$(MAKE) -C admin push
	$(MAKE) -C ide push

test:
	bash test.sh

snapshot:
	git tag $(shell date +%Y.%m%d.%H%M-snapshot)
