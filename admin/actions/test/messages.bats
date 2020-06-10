#!/usr/bin/env bats
load util

@test "messages" {
    post $URL/util/messages a=1
    ckline "validation errors"
    post $URL/util/messages <$H/messages.json
    ckline "id"
    run python3 src/util/cache.py get:message:1234567890123456
    ckline '"fiscal_code": "1234567890123456"'
}
