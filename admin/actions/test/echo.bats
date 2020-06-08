#!/usr/bin/env bats
load util

@test "echo put" {
    run http PUT $URL/util/echo hello=world
    run filter grep -v 'x-'
    run filter grep -v 'accept'
    run filter sed -e 's/[0-9]\{1,5\}/X/g'
    ckdiff <<EOF
{
  "__ow_body": "{\"hello\": \"world\"}",
  "__ow_headers": {
    "content-type": "application/json",
    "host": "X.X.X.X:X",
    "user-agent": "HTTPie/X.X.X",
  },
  "__ow_method": "put",
  "__ow_path": "",
  "__ow_query": ""
}
EOF
}
