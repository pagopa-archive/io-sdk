#!/usr/bin/env bats
load util

@test "echo put" {
    iput $URL/util/echo hello=world
    ckline '"__ow_body": "{\"hello\": \"world\"}"'
    ckline '"content-type": "application/json"'
    ckline '"__ow_method": "put"'
    ckline '"__ow_path": ""'
    ckline '"__ow_query": ""'
}

@test "echo get" {
    get "$URL"'/util/echo/extra/path?q=search&type=string'
    ckline '"__ow_body": ""'
    ckline '"__ow_method": "get"'
    ckline '"__ow_path": "/extra/path"'
    ckline '"__ow_query": "q=search&type=string"'
}
