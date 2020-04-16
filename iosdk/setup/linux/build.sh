#!/bin/bash
set -e
cd $(dirname "${BASH_SOURCE[0]}")
VERSION=${1:?version}
mkdir -p deb/iosdk_${VERSION}/usr/local/bin/
mkdir -p deb/iosdk_${VERSION}/DEBIAN/
mkdir -p rpm
sed -e "/Version/cVersion: ${VERSION}" control > deb/iosdk_${VERSION}/DEBIAN/control

cp bin/* deb/iosdk_${VERSION}/usr/local/bin/

pushd deb
dpkg-deb --build iosdk_${VERSION}
popd

pushd rpm
alien -rg --bump 0 ../deb/iosdk_${VERSION}.deb
sed -i '/%dir/d' iosdk-${VERSION}/iosdk-${VERSION}-1.spec
cd iosdk-${VERSION}
rpmbuild --buildroot $PWD -bb iosdk-${VERSION}-1.spec
popd

