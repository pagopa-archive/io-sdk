#!/bin/bash
set -e
export NODENV_VERSION=12.18.0
export PYENV_VERSION=3.7.7
export GOENV_VERSION=1.13.12
export LOCAL="$HOME/.local"
export BIN="$LOCAL/bin"
mkdir -p "$BIN"


case "$(uname)" in
    (Darwin)
        brew install readline xz zip unzip
        WSK_INSTALL=https://github.com/apache/openwhisk-cli/releases/download/1.0.0/OpenWhisk_CLI-1.0.0-mac-amd64.zip
        JQ_INSTALL=https://github.com/stedolan/jq/releases/download/jq-1.6/jq-osx-amd64
        curl -sL $WSK_INSTALL >$BIN/wsk.zip
        unzip -o $BIN/wsk.zip wsk -d $BIN 
    ;;
    (Linux)
        if test -e /etc/os-release
        then source /etc/os-release
             case "$VERSION_CODENAME" in
                 (bionic) sudo apt-get install -y libicu60 libvpx5 ;;
                 (focal) echo -e "8\n41\n" | apt-get install -y libicu-dev libvpx-dev ;;
                 (*) echo "unsupported ubuntu version" ; exit 1 ;;
             esac
        else echo "sorry only ubuntu 18/20 is supported" ; exit 1
        fi
	    sudo apt-get update
        sudo apt-get install -y build-essential libssl-dev zlib1g-dev libbz2-dev \
            libreadline-dev libsqlite3-dev wget curl libncurses5-dev libncursesw5-dev \
            xz-utils  libffi-dev liblzma-dev python-openssl git zip
        # libraries required to run browsers and support playwright
        sudo apt-get install -y gstreamer1.0-libav gstreamer1.0-plugins-bad \
            gstreamer1.0-plugins-base gstreamer1.0-plugins-good libasound2 libatk-bridge2.0-0 \
            libatk1.0-0 libatspi2.0-0 libbrotli1 libcairo-gobject2 libcairo2 libcups2 libdbus-1-3 \
            libdbus-glib-1-2 libdrm2 libegl1 libenchant1c2a libepoxy0 libfontconfig1 libfreetype6 \
            libgbm1 libgdk-pixbuf2.0-0 libgl1 libgles2 libglib2.0-0 libgstreamer-gl1.0-0 \
            libgstreamer1.0-0 libgtk-3-0 libgtk2.0-0 libharfbuzz-icu0 libharfbuzz0b libhyphen0 \
            libjpeg-turbo8 libnotify4 libnspr4 libnss3 libopenjp2-7 libopus0 libpango-1.0-0 \
            libpangocairo-1.0-0 libpangoft2-1.0-0 libpng16-16 libsecret-1-0 libwayland-client0 \
            libwayland-egl1 libwayland-server0 libwebp6 libwebpdemux2 libwoff1 libx11-6 libx11-xcb1 \
            libxcb-dri3-0 libxcb-shm0 libxcb1 libxcomposite1 libxcursor1 libxdamage1 libxext6 libxfixes3 \
            libxi6 libxkbcommon0 libxml2 libxrandr2 libxrender1 libxslt1.1 libxt6 libxtst6
        WSK_INSTALL=https://github.com/apache/openwhisk-cli/releases/download/1.0.0/OpenWhisk_CLI-1.0.0-linux-amd64.tgz
        JQ_INSTALL=https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64
        curl -sL $WSK_INSTALL >$BIN/wsk.tgz
        tar xzvf $BIN/wsk.tgz -C $BIN  wsk 
    ;;
    (*)
        echo unsupported operating system
        exit 1
    ;;
esac
curl -sL $JQ_INSTALL >$BIN/jq
chmod +x $BIN/jq $BIN/wsk

mkdir -p $BIN

# install
export PYENV_ROOT="$HOME/.pyenv"
export NODENV_ROOT="$HOME/.nodenv"
export GOENV_ROOT="$HOME/.goenv"

if ! test -d $GOENV_ROOT
then git clone https://github.com/syndbg/goenv.git $GOENV_ROOT
fi
if ! test -e $NODENV_ROOT
then git clone https://github.com/nodenv/nodenv.git $NODENV_ROOT
     mkdir "$NODENV_ROOT/plugins"
     git clone https://github.com/nodenv/node-build.git "$NODENV_ROOT"/plugins/node-build
fi
if ! test -e $PYENV_ROOT
then git clone https://github.com/pyenv/pyenv.git $PYENV_ROOT
fi
if ! test -e $LOCAL/bats
then git clone https://github.com/sstephenson/bats.git $LOCAL/bats
    "$LOCAL/bats/install.sh" $LOCAL
fi

export PATH="$HOME/.local/bin:$GOENV_ROOT/bin:$NODENV_ROOT/bin:$PYENV_ROOT/bin:$PATH"

eval "$(nodenv init -)"
nodenv install $NODENV_VERSION -s
echo $NODENV_VERSION >.node-version

eval "$(pyenv init -)"
#eval "$(pyenv virtualenv-init -)"
pyenv install $PYENV_VERSION -s
echo $PYENV_VERSION >.python-version

eval "$(goenv init -)"
goenv install $GOENV_VERSION -s
echo $GOENV_VERSION >.go-version

# etc
python3 -mpip install redis==3.4.1 httpie==2.1.0
npm install -g https://apigcp.nimbella.io/downloads/nim/nimbella-cli.tgz

