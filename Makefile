.PHONY: all
all: .version
	$(MAKE) -C iosdk
	$(MAKE) -C admin
	$(MAKE) -C ide

.PHONY: .version
.version:
	@if test -z `git tag --points-at HEAD` ;\
	then git rev-parse --abbrev-ref HEAD >.version ;\
	else git tag --points-at HEAD | head -1 >.version ;\
	fi

snapshot:
	git tag $(shell date +snap-%Y-%m%d-%H%M)
