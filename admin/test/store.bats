#!/usr/bin/env bats
load util

@test "store" {
    post $URL/util/store <"$H/import.json"
    ckline '"message:ISPXNB32R82Y766D": true,'
    run python3 src/util/cache.py "scan:*"
    run wsk action invoke util/cache -r -p scan "*"
    ckline '"message:ISPXNB32R82Y766D"'
}
