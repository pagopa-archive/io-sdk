.PHONY: build release snapshot
build:
	$(MAKE) -C admin
	$(MAKE) -C ide
	$(MAKE) -C iosdk

release:
	test -n "$(VER)"
	$(MAKE) IOSDK_VER=$(VER)
	$(MAKE) IOSDK_VER=$(VER) -C iosdk/setup

snapshot:
	make VER=$(shell date +%Y.%m%d.%H%M-snapshot) release
