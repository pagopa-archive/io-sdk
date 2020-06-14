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
	sudo apt-get update
        sudo apt-get install -y build-essential libssl-dev zlib1g-dev libbz2-dev \
            libreadline-dev libsqlite3-dev wget curl libncurses5-dev libncursesw5-dev \
            xz-utils  libffi-dev liblzma-dev python-openssl git zip
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
