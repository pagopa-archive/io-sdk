#!/usr/bin/env bats
load util

@test "util/import" {
   run http POST $URL/util/import login=demo password=demo
   run filter jq
   ckdiff <$H/import.json
}
