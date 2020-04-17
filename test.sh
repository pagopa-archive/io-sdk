#!/bin/bash
iosdk/iosdk stop
rm -Rvf $HOME/tmp-iosdk-test
iosdk/iosdk init $HOME/tmp-iosdk-test pagopa/io-sdk-javascript --io-apikey=123456890 --wskprops
docker pull library/redis:5
iosdk/iosdk -v start --skip-pull-images --skip-docker-version --skip-open-browser
docker exec iosdk-theia bash /home/project/build.sh
CHECK=ISPXNB32R82Y766F
DATA="${1:-$HOME/tmp-iosdk-test/data/data.xlsx}"
URL="http://localhost:3280/api/v1/web/guest/iosdk/import"
JSON='{"file": "'$(base64 $DATA)'"}'
HEAD="Content-Type: application/json"
if curl -s $URL -H "$HEAD" -d "$JSON"  | grep $CHECK >/dev/null
then echo SUCCESS ; exit 0
else echo FAIL ; exit 1
fi
iosdk/iosdk stop
