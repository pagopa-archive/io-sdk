VER?=$(shell git tag --points-at HEAD | head -1)

.PHONY: preflight branch build test release release_mac snapshot all clean
branch:
	$(MAKE) build 
	$(MAKE) test

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
	-$(MAKE) -C admin clean
	-$(MAKE) -C ide clean
	-$(MAKE) -C iosdk clean
	
build: preflight
	$(MAKE) -C admin
	$(MAKE) -C ide
	$(MAKE) -C iosdk

push:
	docker login
	$(MAKE) -C admin push
	$(MAKE) -C ide push

test:
	# test cli
	make -C iosdk test
	# test execution
	bash test.sh
	# test actions
	make -C admin test

snapshot:
	date +%Y.%m%d.%H%M-snapshot >.snapshot
	git tag "$(shell cat .snapshot)"
	git push origin master --tags
	git tag -d "$(shell cat .snapshot)"

preflight:
	echo "checking required versions"
	node -v | grep v12
	python3 -V | grep 3.7
	go version | grep go1.13
