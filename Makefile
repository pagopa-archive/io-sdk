.PHONY: iosdk
iosdk: 
	go build -o bin/iosdk

.PHONY: test
test:
	cd wskide && go test -v
