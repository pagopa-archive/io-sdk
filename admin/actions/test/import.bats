#!/usr/bin/env bats
load util

@test "util/import" {
   post $URL/util/import login=demo password=demo
   run filter jq
   ckdiff <$H/import.json
}
