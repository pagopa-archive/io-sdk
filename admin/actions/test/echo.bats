#!/usr/bin/env bats
load util

@test "echo put" {
    run http PUT $URL/util/echo hello=world
    run filter gawk '!/x-/ { gsub("[0-9]+", "X") ; print }' 
    ckdiff <<EOF
{
  "__ow_body": "{\"hello\": \"world\"}",
  "__ow_headers": {
    "accept": "application/json, */*",
    "accept-encoding": "gzip, deflate",
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
