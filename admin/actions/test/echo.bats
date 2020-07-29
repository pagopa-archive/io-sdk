#!/usr/bin/env bats
load util

@test "echo put" {
    iput $URL/util/echo hello=world
    run filter grep -v 'x-'
    run filter grep -v 'accept'
    run filter sed -e 's/[0-9]\{1,5\}/X/g' 
    run filter sed -e 's/ //g'
    ckdiff <"$H/echo.json"
}
