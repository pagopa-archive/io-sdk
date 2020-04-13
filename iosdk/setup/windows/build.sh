#!/bin/bash
cd "$(dirname $0)"
sed "s/iosdk-setup.exe/iosdk_$VER.exe/g" <iosdk.nsi.in >iosdk.nsi 
wine '/root/.wine/drive_c/Program Files/NSIS/makensis.exe' iosdk.nsi
rm iosdk.nsi


