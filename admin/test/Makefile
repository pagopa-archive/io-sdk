.PHONY: all bats playwright

all: bats playwright

bats:
	if ! bats bats ; then echo "FAIL bats: check  /tmp/debug.log"; exit 1 ; fi


playwright: playwright/node_modules
	if ! node playwright/testui.js ; then echo "FAIL playwright"; exit 1 ; fi

playwright/node_modules:
	cd playwright && npm install



