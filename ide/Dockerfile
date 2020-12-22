FROM pagopa/theia:2020.06.08
USER root
RUN ARCH= && dpkgArch="$(dpkg --print-architecture)" \
    && case "${dpkgArch##*-}" in \
    amd64) ARCH='x86_64';; \
    ppc64el) ARCH='ppc64le';; \
    s390x) ARCH='s390x';; \
    arm64) ARCH='arm64';; \
    armhf) ARCH='armv7l';; \
    i386) ARCH='x86';; \
    *) echo "unsupported architecture"; exit 1 ;; \
    esac \
    && curl -SLO "https://download.docker.com/linux/static/stable/$ARCH/docker-18.06.3-ce.tgz" \
    && tar -xzf docker-18.06.3-ce.tgz -C /usr/local \
    && ln -s /usr/local/docker/docker /usr/local/bin

RUN groupadd --gid 999 docker
RUN usermod -a -G docker theia

USER theia
ADD ide-plugin /home/theia/plugins/ide-plugin
RUN sudo find /home/theia -uid 0 -exec chown theia:theia {} \;
#RUN cd /home/theia/plugins/ide-plugin && npm install && npm run build

RUN sudo git clone https://github.com/apache/openwhisk-wskdebug.git /home/wskdebug
#RUN cd /home/wskdebug && sudo npm install -g @openwhisk/wskdebug --unsafe-perm=true --allow-root
RUN cd /home/wskdebug && sudo sed -i 's|0.0.0.0:${hostRuntimePort}|172.17.0.1:${hostRuntimePort}|g' src/invoker.js && sudo npm install && sudo ln -s /home/wskdebug/wskdebug.js /usr/local/bin/wskdebug

RUN sudo find /home/theia -uid 0 -exec chown theia:theia {} \;

