#!/usr/bin/env bats
load util

@test "upload" {
    fpost $URL/util/upload name="mike" cv@$H/upload.txt
    ckline '"cv": "SGVsbG8sIHdvcmxkLgo="'
    ckline '"name": "mike"'
}