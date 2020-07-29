#!/usr/bin/env bats
load util

@test "util/import" {

   ipost $URL/util/import
   ckline '"name": "url"'
   ckline '"name": "username"'
   ckline '"name": "password"'
}

@test "util/import post" {
   ipost $URL/util/import url=https://raw.githubusercontent.com/pagopa/io-sdk/master/docs/sample.json
   ckline '"fiscal_code": "AAAAAA00A00A000A"'
   ckline '"markdown":'
   ckline '"subject":'
}


