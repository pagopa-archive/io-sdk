#!/bin/bash
pushd deb
dpkg-deb --build iosdk_0.5
popd

pushd rpm
alien -r ../deb/iosdk_0.5.deb
popd
