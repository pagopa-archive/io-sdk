#!/usr/bin/env bats
load util


@test "send message" {
    run http -f POST  $URL/util/send dest=SNDMSGTEST1234 subject=Hello markdown=World
    ckline '"id": "832703095"'
    run wsk action invoke util/cache -p scan "send:*" -r
    ckline '"send:SNDMSGTEST1234"'
    run wsk action invoke util/cache -p get "send:SNDMSGTEST1234" -r
    ckdiff <<EOF
{
    "send:SNDMSGTEST1234": {
        "markdown": "World",
        "subject": "Hello"
    }
}
EOF
}