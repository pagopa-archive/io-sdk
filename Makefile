.PHONY: all
build:
	$(MAKE) -C admin
	$(MAKE) -C ide
	$(MAKE) -C iosdk

setup:
	$(MAKE) -C iosdk/setup

snapshot:
	-rm iosdk/iosdk
	git commit -a
	git tag $(shell date +%Y.%m%d.%H%M-snapshot)
	$(MAKE)
