.PHONY: all
all:
	$(MAKE) -C iosdk
	$(MAKE) -C admin
	$(MAKE) -C ide

snapshot:
	-rm iosdk/iosdk
	git commit -a
	git tag $(shell date +snap-%Y-%m%d-%H%M)
	$(MAKE)
