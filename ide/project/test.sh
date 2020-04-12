#!/bin/bash
cd "$(dirname $0)"
AUTH='23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP'
wsk property set \
    --apihost 'http://openwhisk:3280' \
    --auth "$AUTH"
wsk action update hello index.js
wsk action invoke hello -r >/tmp/result.txt
wsk action delete hello
if grep "Hello, world" /tmp/result.txt
then echo "SUCCESS" ; exit 0
else echo "FAILED" ; exit 1
fi
