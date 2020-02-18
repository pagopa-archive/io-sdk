.PHONY: iosdk
iosdk: 
	go build -o bin/iosdk
.PHONY: pkger
pkger:
	pkger
	go build -o bin/iosdk
