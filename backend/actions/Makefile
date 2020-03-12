
.PHONY: deploy
deploy: zip/send.zip
	wsk project deploy

src/send.zip: src/send/main.go src/send/const.go src/send/send.go
	zip -j $@ $^

zip/send.zip: src/send.zip
	docker run -i openwhisk/action-golang-v1.11 -compile main <$< >$@  
 