FROM ubuntu:18.04
ENV NODE_INSTALL=https://deb.nodesource.com/setup_10.x
ENV WSK_INSTALL=https://github.com/apache/openwhisk-cli/releases/download/1.0.0/OpenWhisk_CLI-1.0.0-linux-amd64.tgz
ADD project /home/project
ADD ide-plugin /home/plugins/ide-plugin
RUN apt-get update &&\
   apt-get install -y curl sudo zip jq git &&\
   curl -sL ${NODE_INSTALL}| sudo -E bash - &&\
   apt-get install -y nodejs gcc g++ make &&\
   curl -sL ${WSK_INSTALL} | tar xzvf - -C /usr/bin wsk &&\
   groupadd theia ; useradd theia -g theia -m;\
   chown -R theia:theia /home/project /home/plugins &&\
   echo "theia ALL=(ALL) NOPASSWD: ALL" >>/etc/sudoers &&\
   npm install -g yarn typescript jest
USER theia
RUN cd /home/plugins/ide-plugin && npm install && npm run build
WORKDIR /home/theia
ADD latest.package.json /home/theia/package.json
RUN yarn --pure-lockfile && \
    NODE_OPTIONS="--max_old_space_size=4096" yarn theia build && \
    yarn --production && \
    yarn autoclean --init && \
    echo *.ts >> .yarnclean && \
    echo *.ts.map >> .yarnclean && \
    echo *.spec.* >> .yarnclean && \
    yarn autoclean --force && \
    yarn cache clean
RUN sudo dpkg --purge gcc g++ make python
EXPOSE 3000
USER theia
ENV HOME /home/theia
ENV SHELL /bin/bash
ENV USE_LOCAL_GIT true
ENV THEIA_DEFAULT_PLUGINS=local-dir://home/plugins
ENTRYPOINT [ "node", "/home/theia/src-gen/backend/main.js", "/home/project", "--hostname=0.0.0.0" ]
