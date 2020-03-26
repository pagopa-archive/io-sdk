#!/usr/bin/env bats
load util

@test "send message" {
    run http -f POST  $URL/util/send fiscal_code=SNDMSGTEST1234 subject=Hello markdown=World
    ckline '"id": "832703095"'
    run wsk action invoke util/cache -p scan "sent:*" -r
    ckline '"sent:SNDMSGTEST1234"'
    run wsk action invoke util/cache -p get "sent:SNDMSGTEST1234" -r
    ckdiff <<EOF
{
    "sent:SNDMSGTEST1234": {
        "markdown": "World",
        "subject": "Hello"
    }
}
EOF
}

@test "send message real" {
    skip
    run 
    http -f POST $URL/iosdk/send fiscal_code=ISPXNB32R82Y766F subject="Hello World" markdown="This must be a long string of at least 80 characters, so let me talk a bit more to get that length."
    ckline id

}