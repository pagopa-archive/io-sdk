cd /tmp
virtualenv virtualenv
source virtualenv/bin/activate
pip3 install redis
apt-get update && apt-get -y install zip
zip -q -r /mnt/virtualenv.zip virtualenv
