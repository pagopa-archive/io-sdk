#!/bin/bash
echo "You have an hour to debug this build"
echo 'do "touch /tmp/continue" to stop the wait'
for i in $(seq 1 60) 
do echo $i
    sleep 60
    if test -e /tmp/continue
    then break
    fi
done
