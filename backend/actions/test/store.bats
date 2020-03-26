#!/usr/bin/env bats
load util

@test "store" {
    run http POST $URL/util/store <$H/import.json
    ckline '"message:ISPXNB32R82Y766D": true,'
    run python3 deploy/util/cache.py "scan:*"
    ckline '"message:ISPXNB32R82Y766D"'
}
