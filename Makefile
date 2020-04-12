.PHONY: all
all:
	$(MAKE) -C iosdk
	$(MAKE) -C admin
	$(MAKE) -C ide

snapshot:
	git tag $(shell date +snap-%Y-%m%d-%H%M)
