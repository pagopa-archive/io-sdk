#!/usr/bin/env bats
load util

@test "messages" {
    ipost $URL/util/messages a=1
    ckline "validation errors"
    post $URL/util/messages <"$H/messages.json"
    ckline "id"
    run wsk action invoke util/cache -r -p get message:1234567890123456
    ckline '"fiscal_code": "1234567890123456"'
}
