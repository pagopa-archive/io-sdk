#!/usr/bin/env bats
load util

@test "store" {
    http PUT $URL/util/store <$H/import.json
    ckline "validation errors"
    run http POST $URL/util/messages <$H/messages.json
    ckline "id"
    run python3 deploy/util/cache.py get:1234567890123456
    ckline '"fiscal_code": "1234567890123456"'
}
