FROM alpine:3.10 as build
ENV GO111MODULE=on
RUN apk update ; apk add go git busybox-extras alpine-sdk ;\
git clone https://github.com/sciabarracom/openwhisk-cli ;\
cd openwhisk-cli ;\
go build -o /usr/bin/wsk 
FROM actionloop/ide-js-base:latest
USER root
ENV DOCKER_VERSION=18.06.3-ce
COPY --from=build /usr/bin/wsk /usr/bin/wsk
RUN apk add curl sudo zip jq ;\
    echo "theia ALL=(ALL) NOPASSWD: ALL" >>/etc/sudoers ;\
    npm -g install jest typescript ;\
    curl -sL \
  https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz \
  | tar xzvf -  -C /usr/bin --strip 1 docker/docker
