FROM i386/debian:buster
ENV NSIS=https://sourceforge.net/projects/nsis/files/NSIS%203/3.05/nsis-3.05-setup.exe/download
RUN apt-get update &&\
    apt-get install -y wine curl ruby ruby-dev rubygems build-essential rpm &&\
    gem install --no-document fpm &&\
    curl -sL ${NSIS} >/tmp/nsis.exe &&\
    wine /tmp/nsis.exe /S
