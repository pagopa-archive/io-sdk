.PHONY: iosdk
iosdk: 
	env GO111MODULE=on go build -o bin/iosdk

.PHONY: test
test:
	cd wskide && go test -v
