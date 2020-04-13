#!/bin/bash
if test -z `git tag --points-at HEAD` 
then git rev-parse --abbrev-ref HEAD
else date +%Y.%m%d.%H%M-snapshot
fi
