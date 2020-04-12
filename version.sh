#!/bin/bash
if test -z "$IOSDK_VERSION"
then 
    if test -z `git tag --points-at HEAD` 
    then git rev-parse --abbrev-ref HEAD
    else git tag --points-at HEAD | head -1 >.version
    fi
else echo "$IOSDK_VERSION"
fi
