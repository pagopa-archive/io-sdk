#!/usr/bin/env bats
load util

@test "send message" {
    fpost $URL/util/send fiscal_code=SNDMSGTEST1234 subject=Hello markdown=World
    ckline '"id": "832703095"'
    run wsk action invoke util/cache -p scan "sent:*" -r
    ckline '"sent:SNDMSGTEST1234"'
    run wsk action invoke util/cache -p get "sent:SNDMSGTEST1234" -r
    ckdiff <<EOF
{
    "sent:SNDMSGTEST1234": {
        "fiscal_code": "SNDMSGTEST1234",
        "markdown": "World",
        "subject": "Hello"
    }
}
EOF
}

