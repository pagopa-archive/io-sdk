#!/bin/bash
set -e
iosdk/iosdk stop
sudo rm -Rvf $HOME/tmp-iosdk-test
docker pull library/redis:5
echo "****** INIT"
iosdk/iosdk init $HOME/tmp-iosdk-test pagopa/io-sdk-javascript --io-apikey=123456890 --wskprops
echo "****** START"
iosdk/iosdk -v start --skip-pull-images --skip-docker-version --skip-open-browser
echo "****** BUILD"
docker exec iosdk-theia bash -c 'sudo bash /home/project/build.sh'
echo "****** STATUS"
iosdk/iosdk status
CHECK=ISPXNB32R82Y766F
DATA="${1:-$HOME/tmp-iosdk-test/data/data.xlsx}"
URL="http://localhost:3280/api/v1/web/guest/iosdk/import"
JSON='{"file": "'$(base64 $DATA | tr -d '\n')'"}'
HEAD="Content-Type: application/json"
if curl -s $URL -H "$HEAD" -d "$JSON"  | grep $CHECK >/dev/null
then echo SUCCESS ; exit 0
else echo FAIL ; exit 1
fi
iosdk/iosdk stop
