.PHONY: iosdk
iosdk: 
	go build -o bin/iosdk

.PHONY: test
test:
	cd wskide && go test -v


.PHONY: linux-packages
linux-packages:
	cd installer/linux && bash -x create-packages.sh
