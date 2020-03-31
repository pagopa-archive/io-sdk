.PHONY: iosdk
iosdk: 
	go build -o bin/iosdk

deploy:
	$(MAKE) -C backend deploy

.PHONY: test
test:
	cd wskide && go test -v
