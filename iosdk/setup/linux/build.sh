#!/bin/bash
cd /mnt
fpm -s dir -t deb -n iosdk -v $VER ./bin=/usr
fpm -s dir -t rpm -n iosdk -v $VER ./bin=/usr
