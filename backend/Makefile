IMG=actionloop/iosdk:latest
<<<<<<< HEAD
=======
AUTH=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP
>>>>>>> e6b2df6053b8c531571140e1242f37e38454cd8b

.PHONE: build
build:
	node -v | grep v10
	npm install
	npm run build
	docker build -t $(IMG) .

.PHONY: start
start:
	docker run -ti \
	-p 3280:3280 -p 3281:3281 \
	--rm --name openwhisk --hostname openwhisk \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-e CONTAINER_EXTRA_ENV=__OW_DEBUG_PORT=8081 \
	$(IMG)

.PHONY: stop
stop:
	-docker exec openwhisk stop

AUTH=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP

.PHONY: wsk
wsk:
	wsk property set --apihost http://localhost:3280 --auth $(AUTH)

.PHONY: deploy
deploy: wsk
	cd actions ; wsk project deploy 