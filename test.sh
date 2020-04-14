#!/bin/bash
iosdk/iosdk stop
rm -Rvf $HOME/tmp-iosdk-test
iosdk/iosdk init $HOME/tmp-iosdk-test pagopa/io-sdk-javascript --io-apikey=123456890 --wskprops
docker pull library/redis:5
iosdk/iosdk start --do-not-pull-images
docker exec -ti iosdk-theia bash /home/project/build.sh
CHECK=ISPXNB32R82Y766F
DATA="${1:-$HOME/tmp-iosdk-test/data/data.xlsx}"
URL="$(wsk action get iosdk/import --url | tail +2)"
JSON="$(jq -r -n --arg file "$(base64 $DATA)" '{file: $file}' | tr -d '[\n\t ]')"
HEAD="Content-Type: application/json"
if curl -s $URL -H "$HEAD" -d "$JSON"  | grep $CHECK >/dev/null
then echo SUCCESS ; exit 0
else echo FAIL ; exit 1
fi
