#!/bin/bash
set -e
cd $(dirname "${BASH_SOURCE[0]}")
ls -la
cp ../../bin/iosdk deb/iosdk_0.5/usr/local/bin/

pushd deb
dpkg-deb --build iosdk_0.5
popd

pushd rpm
alien -r ../deb/iosdk_0.5.deb
popd
