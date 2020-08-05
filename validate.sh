#!/bin/bash
VERSION="${1:?usage: <version>}"
set -e
# install
PKG=IO-SDK-macos-installer-x64-${VERSION}.pkg
URL=https://github.com/pagopa/io-sdk/releases/download/${VERSION}/$PKG
echo downloading $URL
curl -sL $URL >/tmp/$PKG
sudo installer -pkg /tmp/$PKG -target /
# check version
V=$(iosdk --version 2>&1)
if test "$V" != "$VERSION"
then echo $V: wrong version, expected $VERSION
     echo FAIL
     exit 1
fi
# init
rm -Rvf $HOME/tmp-iosdk-validation
/usr/local/bin/iosdk init $HOME/tmp-iosdk-validation pagopa/io-sdk-javascript --io-apikey=123456890
# test
/usr/local/bin/iosdk stop
/usr/local/bin/iosdk -v start --skip-open-browser
docker exec iosdk-theia bash -c 'bash /home/project/build.sh'
/usr/local/bin/iosdk status
CHECK=ISPXNB32R82Y766F
DATA="${1:-$HOME/tmp-iosdk-validation/data/data.xlsx}"
URL="http://localhost:3280/api/v1/web/guest/iosdk/import"
JSON='{"file": "'$(base64 $DATA | tr -d '\n')'"}'
HEAD="Content-Type: application/json"
curl -s $URL -H "$HEAD" -d "$JSON"  | grep $CHECK >/dev/null
RES=$?
# cleanup
rm -Rvf $HOME/tmp-iosdk-validation
/usr/local/bin/iosdk stop
yes | sudo /Library/IO-SDK/$VERSION/uninstall.sh
# result
if test "$RES" == "0" 
then echo SUCCESS ; exit 0
else echo FAIL ; exit 1
fi
