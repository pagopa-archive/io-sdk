.PHONY: all bats playwright

test: bats playwright

bats:
	bats bats

playwright: playwright/node_modules
	cd playwright && npm test 

playwright/node_modules:
	cd playwright && npm install

