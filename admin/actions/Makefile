.PHONY: build
build: deploy/send.zip deploy/sendpy.zip deploy/cache.zip deploy/messages.zip deploy/store.zip

.PHONY: deploy
deploy: build 
	wsk project deploy

.PHONY: undeploy
undeploy: setup
	wsk project undeploy

.PHONY: test
test: deploy
	# warm up 
	python3 -mpip install redis==3.4.1 httpie==2.1.0
	-wsk action list |\
	awk '/guest/ { print $$1}' |\
	xargs -L1 wsk action invoke -r
	# test
	if ! bats test ; then cat /tmp/debug.log ; exit 1 ; fi
	

deploy/send.zip: $(filter-out  %_test.go, $(wildcard src/send/*.go))
	zip - -r -j $^ | docker run -i openwhisk/action-golang-v1.11 -compile main >$@

virtualenv.zip:
	docker run -i --entrypoint=/bin/bash -v $(shell pwd):/mnt \
	openwhisk/actionloop-python-v3.7 /mnt/venv.sh

deploy/cache.zip: src/util/cache.py virtualenv.zip
	cp virtualenv.zip $@
	cp $< __main__.py
	zip -u $@ __main__.py

deploy/store.zip: src/util/store.py virtualenv.zip
	cp virtualenv.zip $@
	cp $< __main__.py
	zip -u $@ __main__.py

deploy/messages.zip: src/util/messages.py virtualenv.zip
	cp virtualenv.zip $@
	cp $< __main__.py
	zip -u $@ __main__.py

deploy/sendpy.zip: src/util/send.py virtualenv.zip
	cp virtualenv.zip $@
	cp $< __main__.py
	zip -u $@ __main__.py

.PHONY: clean
clean:
	-rm deploy/*.zip virtualenv.zip __main__.py
