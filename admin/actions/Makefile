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
	python3 -m pip install redis
	- wsk action invoke util/echo -r >/dev/null
	- bats test/cache.bats >/dev/null
	bats test

deploy/send.zip: $(filter-out  %_test.go, $(wildcard src/send/*.go))
	zip - -r -j $^ | docker run -i openwhisk/action-golang-v1.11 -compile main >$@

virtualenv.zip:
	python3 -m virtualenv virtualenv
	virtualenv/bin/pip3 install redis
	zip -r $@ virtualenv

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
	-rm deploy/*.zip virtualenv.zip

