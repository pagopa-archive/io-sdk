
.PHONY: deploy
deploy: zip/send.zip
	wsk project deploy

zip/send.zip: src/send/main.go src/send/const.go
	zip -j $@ $^
